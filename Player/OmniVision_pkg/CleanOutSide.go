package OmniVision_pkg
import (
	"image"
	"image/color"
	"gocv.io/x/gocv"
)
// Clean_OutField is a function that clean all the noise outside the footbal field using last robot position.
// Is was consider outside of the field when is 1 meter way of outside lines (10 pixels)
//
// Parameters:
//
//   - `readed_field` : gocv.Mat -> Matrix 160x160  to read and save real field
//   - `x` : int -> previous x robot's coordinate
//   - `y` : int -> previous y robot's coordinate
//
// Outputs:
//
//   - `readed_field` : gocv.Mat -> Binary Image of lines in real distance with field's outside with no noise

func (vision *OmniVision)Clean_OutField(readed_field *gocv.Mat, x int, y int) {
	black := color.RGBA{0, 0, 0, 0}
	Limit_L_x := 80 - ((B / 20) - y/10) - 10
	Limit_R_x := 80 + ((B / 20) + y/10) + 10
	Limit_T_y := 80 - ((A / 20) - x/10) - 10
	Limit_B_y := 80 + ((A / 20) + x/10) + 10
	
	if(Limit_L_x>0 && Limit_L_x <160){
		gocv.Rectangle(readed_field, image.Rect(0, 0, Limit_L_x, 160), black, -1)
	}
	if(Limit_R_x>0 && Limit_R_x <160){
		gocv.Rectangle(readed_field, image.Rect(Limit_R_x, 0, 160, 160), black, -1)
	}
	if(Limit_T_y>0 && Limit_T_y <160){
		gocv.Rectangle(readed_field, image.Rect(0, 0, 160, Limit_T_y), black, -1)
	}
	if(Limit_B_y>0 && Limit_B_y <160){
		gocv.Rectangle(readed_field, image.Rect(0, Limit_B_y, 160, 160), black, -1)
	}
}
