package OmniVision_pkg

import (
	"fmt"
	"image"
	GK "player/GK_Position"
	"gocv.io/x/gocv"
)

// The Global_Localise function obtains the robot's position and angle based on the entire field.
//  It is a method of the OmniVision struct.
//
// Returns:
//
//   - `x` : int ->  Robot X Position in centimetres
//   - `y` : int ->  Robot Y Position in centimetres
//   - `angle` : int ->  Robot angle
//
func (vision *OmniVision) Global_Localize() (x, y, angle int) {
	// Variables
	result := gocv.NewMat()
	defer result.Close()
	vision.ideal_field.CopyTo(&vision.ideal_field_to_edit)
	vision.calc_field.SetTo(gocv.NewScalar(0, 0, 0, 0))

	// # HSV Filter
	gocv.CvtColor(vision.original_image, &vision.hsv_image, gocv.ColorBGRToHSV)
	gocv.InRangeWithScalar(vision.hsv_image, vision.lowerMask, vision.upperMask, &vision.viewed_field)
	
	// # Get Real Field
	GetRealField(&vision.viewed_field, &vision.calc_field)
	
	// # Rotate and Clean Outside Lines
	vision.angle  = vision.adjustAngle(&vision.calc_field)

	// # Template Matching
	gocv.MatchTemplate(vision.ideal_field_to_edit, vision.calc_field, &result, gocv.TmCcoeff, vision.mask) 
	_, _, _, maxLoc := gocv.MinMaxLoc(result)
	
	// Absolute Coordenates
	vision.x_vision, vision.y_vision = Pixels_to_MM(maxLoc.X+80, maxLoc.Y+80)
	fmt.Println("Vision Localization X=", vision.x_vision, " Y=", vision.x_vision)
	
return vision.x_vision, vision.y_vision, vision.angle
}

// The Close_Localize function obtains the robot's position and angle based on the previous position,
// i.e. it only has a radius of 1 metre (10 pixels) to move from the last position.
//  It is a method of the OmniVision struct.
//
// Returns:
//
//   - `x` : int ->  Robot X Position in centimetres
//   - `y` : int ->  Robot Y Position in centimetres
//   - `angle` : int ->  Robot angle
//
func (vision *OmniVision) Close_Localize() (x float64, y float64, angle int) {
	// Variables
	vision.ideal_field.CopyTo(&vision.ideal_field_to_edit)
	x_vision, y_vision := CM_to_Pixels(vision.x, vision.y)
	vision.calc_field.SetTo(gocv.NewScalar(0, 0, 0, 0))

	// # HSV Filter
	gocv.CvtColor(vision.original_image, &vision.hsv_image, gocv.ColorBGRToHSV)
	gocv.InRangeWithScalar(vision.hsv_image, vision.lowerMask, vision.upperMask, &vision.viewed_field)
	
	// # Get Real Field
	GetRealField(&vision.viewed_field, &vision.calc_field)
	
	// # Rotate and Clean Out
	vision.adjustAngle(&vision.calc_field)
		
	// # Region of Intrest to Template matching
	ideal_field_roi := vision.ideal_field_to_edit.Region(image.Rect(x_vision-90, y_vision-90, x_vision+90, y_vision+90))
	
	// # Template Matching
	gocv.MatchTemplate(ideal_field_roi, vision.calc_field, &vision.result, gocv.TmCcoeff, vision.mask) 
	_, _, _, maxLoc_roi := gocv.MinMaxLoc(vision.result)

	// Absolute Coordenates
	vision.x_vision, vision.y_vision = Pixels_to_MM(x_vision-90+ maxLoc_roi.X+80, y_vision-90+maxLoc_roi.Y+80)
		
        // # Kalman Filter
	if vision.ID == 0 {
		x_lidar, y_lidar, angle := GK.Get_GKPosition(A / 2)
		if x != 999999 || y != 999999 {
			vision.x_vision = x_lidar
			vision.y_vision = y_lidar
			vision.x, vision.y = vision.Update_Kalman()
			if angle > 70 || angle < -70 {
				vision.angle = int(angle)
			}
		} else {
			vision.x, vision.y = vision.Update_Kalman()
		}
	} else {
		vision.x, vision.y = vision.Update_Kalman()
	}
return vision.x, vision.y, vision.angle
}


