package OmniVision_pkg

import (
	"math"
	coms "player/communication"
	pb "player/pb"
)

// Object_in to know if the object is inside of the field or not.
// It is a method of the OmniVision struct.
//
// Parameters:
//
//   - `dist` : float64 -> Distance to object
//   - `angle` : float64 -> Angle
//
// Returns:
//
//   - `In` : bool -> In/Out
//
func (vision *OmniVision) Object_in(dist int, angle int) bool {

	Obj_angle := -vision.angle + angle //angle robot
	X_obj_rel := float64(dist) * math.Sin(float64(Obj_angle)*(math.Pi/180))
	Y_obj_rel := float64(dist) * math.Cos(float64(Obj_angle)*(math.Pi/180))
	X_obj := int(vision.x) + int(X_obj_rel)
	Y_obj := int(vision.y) + int(Y_obj_rel)

	if X_obj > (A/2+25) || X_obj < -(A/2-25) || Y_obj > (B/2+25) || Y_obj < -(B/2-25) {
		return false
	} else {
		return true
	}

}

// SaveOmniData is a function that save all the object in Local Database.
// It is a method of the OmniVision struct.
//
// Parameters:
//
//   - `resp` : *Response_Omni -> YOLO Data (via gRPC)
//
func (vision *OmniVision)SaveOmniData( resp *pb.Response_Omni) {
	var last_ball coms.Ball_st
	var robots_t []coms.Robot_st
	var robots_o []coms.Robot_st
	var posts []coms.Robot_st
	var my_localization coms.Robot_st
	
	my_localization.Coords.X = vision.x
	my_localization.Coords.Y = vision.y
	my_localization.Orientation = vision.angle
	my_localization.Distance = 0
	robots_t = append(robots_t, my_localization)
	first_ball := true
	for _, object := range resp.Objects {
		switch object.Id {
		case 0: // BALL

			if first_ball {
				if vision.Object_in(real_coordenates_omni[object.X][object.Y][2], real_coordenates_omni[object.X][object.Y][3]) {
					var my_ball coms.Ball_st
					my_ball.Coords.X = float64(real_coordenates_omni[object.X][object.Y][0])
					my_ball.Coords.Y = float64(real_coordenates_omni[object.X][object.Y][1])
					my_ball.Angle = real_coordenates_omni[object.X][object.Y][3]
					my_ball.Dist = real_coordenates_omni[object.X][object.Y][2]
					my_ball.Conf = int(object.Conf)
					last_ball = my_ball
					first_ball = false
				}
			}
			break
		case 1: // BLUE_SHIRT
			if vision.Object_in(real_coordenates_omni[object.X][object.Y][2], real_coordenates_omni[object.X][object.Y][3]) {
				var blue_shirt coms.Robot_st
				blue_shirt.Coords.X = float64(real_coordenates_omni[object.X][object.Y][0])
				blue_shirt.Coords.Y = float64(real_coordenates_omni[object.X][object.Y][1])
				blue_shirt.Distance = real_coordenates_omni[object.X][object.Y][2]
				blue_shirt.Conf = int(object.Conf)
				blue_shirt.Angle  = real_coordenates_omni[object.X][object.Y][3]
				if vision.shirt { // if blue
					robots_t = append(robots_t, blue_shirt)
				} else {
					robots_o = append(robots_o, blue_shirt)
				}
			}
			break

		case 2: // GOALPOSTS
			if vision.Object_in(real_coordenates_omni[object.X][object.Y][2], real_coordenates_omni[object.X][object.Y][3]) {
				var goalpost coms.Robot_st
				goalpost.Coords.X = float64(real_coordenates_omni[object.X][object.Y][0])
				goalpost.Coords.Y = float64(real_coordenates_omni[object.X][object.Y][1])
				goalpost.Distance = real_coordenates_omni[object.X][object.Y][2]
				goalpost.Conf = int(object.Conf)
				goalpost.Angle = real_coordenates_omni[object.X][object.Y][3]
				posts = append(posts, goalpost)
			}
			break

		case 3: // PERSON
			if vision.Object_in(real_coordenates_omni[object.X][object.Y][2], real_coordenates_omni[object.X][object.Y][3]) {
				var red_shirt coms.Robot_st
				red_shirt.Coords.X = float64(real_coordenates_omni[object.X][object.Y][0])
				red_shirt.Coords.Y = float64(real_coordenates_omni[object.X][object.Y][1])
				red_shirt.Distance = real_coordenates_omni[object.X][object.Y][2]
				red_shirt.Conf = int(object.Conf)
				red_shirt.Angle = real_coordenates_omni[object.X][object.Y][3]
				robots_o = append(robots_o, red_shirt)
			}
			break

		case 4: // RED_SHIRT
			if vision.Object_in(real_coordenates_omni[object.X][object.Y][2], real_coordenates_omni[object.X][object.Y][3]) {
				var red_shirt coms.Robot_st
				red_shirt.Coords.X = float64(real_coordenates_omni[object.X][object.Y][0])
				red_shirt.Coords.Y = float64(real_coordenates_omni[object.X][object.Y][1])
				red_shirt.Distance = real_coordenates_omni[object.X][object.Y][2]
				red_shirt.Conf = int(object.Conf)
				red_shirt.Angle = real_coordenates_omni[object.X][object.Y][3]
				if vision.shirt { // if blue
					robots_o = append(robots_o, red_shirt)
				} else {
					robots_t = append(robots_t, red_shirt)
				}
			}
			break

		case 5: // ROBOT
			if vision.Object_in(real_coordenates_omni[object.X][object.Y][2], real_coordenates_omni[object.X][object.Y][3]) {
				var robot coms.Robot_st
				robot.Coords.X = float64(real_coordenates_omni[object.X][object.Y][0])
				robot.Coords.Y = float64(real_coordenates_omni[object.X][object.Y][1])
				robot.Distance = real_coordenates_omni[object.X][object.Y][2]
				robot.Angle = real_coordenates_omni[object.X][object.Y][3]
				robots_o = append(robots_o, robot)
			}
			break

		}
	}
	coms.SetBallPosition(last_ball)
	coms.SetRobotsPositions(robots_t, robots_o)
	//coms.SetPostsPosition(posts)
}
