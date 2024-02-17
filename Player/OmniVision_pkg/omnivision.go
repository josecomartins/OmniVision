// This package obtains the position of the robot itself and all the other elements on a robotic soccer field.
//
// The aim is to obtain the image from a catadioptric camera and the data from YOLO via gRPC, 
// obtain the robot's current position and save the resulting data in the Local Database.
//
package OmniVision_pkg

import (
	"context"
	"fmt"
	"image"
	"log"
	coms "player/communication" // Communication pkg
	pb "player/pb" // gRPC Grenerated pkg
	"gocv.io/x/gocv"
	"google.golang.org/grpc"
)

// OmniVision Struct contains all the data needed by OmniVision's software, eliminating the need to create new variables every cycle.
//
type OmniVision struct {
	ID                  int      	// Robot ID
	shirt               bool	// Robot Shirt
	x_vision            int     	// Robot X Position Computer Vision
	y_vision            int     	// Robot Y Position Computer Vision
	x_enc               int     	// Robot X Position Encoders
	y_enc               int		// Robot Y Position Encoders
	x                   float64     // Robot X Position Final
	y                   float64     // Robot Y Position Final
	angle               int     	// Robot Angle
	original_image      gocv.Mat 	// RGB Frame
	hsv_image           gocv.Mat 	// HSV Frame
	viewed_field        gocv.Mat	// Lines Filtered
	calc_field          gocv.Mat	// Real Distance Lines
	calc_rotated_field  gocv.Mat	// Orientated Real Distance Lines
	ideal_field         gocv.Mat 	// Image of Full Field
	ideal_field_to_edit gocv.Mat	// Image of Full Field (for debug)
	mask                gocv.Mat	// TemplateMatching Mask
	result              gocv.Mat	// TemplateMatching Result
	maxLoc              image.Point // TemplateMatching Result Position
	kalman              *Kalman 	// Kalman Variables
	lowerMask           gocv.Scalar	// Low Margin HSV Filter
	upperMask           gocv.Scalar	// Upper Margin HSV Filter
}

// This function init all variables of OmniVision.
// It is a method of the OmniVision struct.
//
func (vision *OmniVision) Init() {
	//Init Variables
	vision.hsv_image = gocv.NewMat()
	vision.calc_field = gocv.NewMatWithSize(160, 160, gocv.MatTypeCV8UC1)
	vision.calc_rotated_field = gocv.NewMatWithSize(160, 160, gocv.MatTypeCV8UC1)
	vision.viewed_field = gocv.NewMat()
	vision.ideal_field = gocv.IMRead("OmniVision_pkg/FIELD.png", gocv.IMReadGrayScale)
	vision.ideal_field_to_edit = gocv.IMRead("OmniVision_pkg/FIELD.png", gocv.IMReadGrayScale)
	white := gocv.NewScalar(255, 255, 255, 255)
	vision.mask = gocv.NewMatWithSizeFromScalar(white, 160, 160, gocv.MatTypeCV8UC1)
	vision.lowerMask = gocv.NewScalar(MIN_H_lines_omni, MIN_S_lines_omni, MIN_V_lines_omni, 0.0)
	vision.upperMask = gocv.NewScalar(MAX_H_lines_omni, MAX_S_lines_omni, MAX_V_lines_omni, 0.0)
	vision.result = gocv.NewMat()
}

// The Run function is responsible for starting all of OmniVision and is in a permanent cycle. 
// In this permanent cycle, it must receive the image and data from YOLO, 
// perform Close_Localization and save the resultant data in the Local Database.
//
// Parameters:
//
//   - `shirt` : bool ->  Robot Shirt (blue=true red=false)
//   - `ip_grpc` : string -> gRPC IP
//   - `ID` : int ->  Robot Number (0 is GK - 1,2,3,4 are field players)
//
func Run(shirt bool, ip_grpc string, ID int) {
	
	fmt.Println("OmniVision!")
	// # Init variables
	var omnivision OmniVision
	omnivision.Init()
	buttons := 0
	side := false
	omnivision.shirt = shirt 
	omnivision.ID = ID
	// Robot Buttons, Robot Shirt, Field Side
	coms.GetButtons(&buttons, &omnivision.shirt, &side) 
	
	// # Init gRPC
	addr := ip_grpc + ":40000"
	fmt.Println("GRPC Omni service openning...", addr)
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewYolo_OmniClient(conn)
	req := pb.Request_Omni_Calib{
		Check: false,
	}
	stream,_:=client.Send_Omni(context.Background(), &req)
	fmt.Println("GRPC Omni service opened!", addr)

	
	// # Global Localization
	fmt.Println("Global localization!") 
	N_GLOBAL_ITERATIONS := 10
	x_global := 0
	y_global := 0
	num_it := 0
	
	for i := 0; i < N_GLOBAL_ITERATIONS; i++ {
		// # Request Data (Image)
		resp, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		omnivision.original_image, _ = gocv.NewMatFromBytes(480, 480, gocv.MatTypeCV8UC3, resp.Omni)
		
		// # Global Localization
		omnivision.Global_Localize()
		if (omnivision.x_vision < 0 && !side) || (omnivision.x_vision > 0 && side) {
			x_global += omnivision.x_vision
			y_global += omnivision.y_vision
			fmt.Println("🌍️ GlobaL Localization X=", omnivision.x_vision, " Y=", omnivision.y_vision)
			num_it++
		}
	}
	if num_it != 0 {
		omnivision.x_vision = x_global / num_it
		omnivision.y_vision = y_global / num_it
	}
	fmt.Println("Localization Ready! X=", omnivision.x_vision, " Y=", omnivision.y_vision)

	// # Init Kalman
	omnivision.Init_Kalman(float64(omnivision.x_vision), float64(omnivision.y_vision))
	
	fmt.Println("OmniVision Ready!")
	for {
		coms.GetButtons(&buttons, &omnivision.shirt, &side)
		// # Relocation
		if buttons == 2 {
			omnivision.Global_Localize()
			omnivision.Reset_Kalman()
		}
		// # Recieve Data (Image and YOLO data)
		resp, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		// Convert gRPC bytes in gocv.Mat
		omnivision.original_image, _ = gocv.NewMatFromBytes(480, 480, gocv.MatTypeCV8UC3, resp.Omni)
		
		// # Close Self-Localization
		omnivision.Close_Localize()
		
		// # Save Data
		omnivision.SaveOmniData(resp)
	}
}
