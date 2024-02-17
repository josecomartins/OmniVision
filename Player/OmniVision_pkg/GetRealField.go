package OmniVision_pkg
import (
	"math"
	"gocv.io/x/gocv"
)
const MIN_LINE_WIDTH = 1
const MAX_LINE_WIDTH = 60

// GetRealField is a function that takes the filtred field and transform it in real field
// The resolution is 10cm for pixel
// Considering that robots can view 6 meter in each direction
//
// Parameters:
//
//   - `image_field` : gocv.Mat ->  Binary Image of Field
//   - `readed_field` : gocv.Mat -> Matrix 120x120 to save real field
//
func GetRealField(image_field *gocv.Mat, readed_field *gocv.Mat) {

	image_field_ptr, _ := image_field.DataPtrUint8()
	readed_field_ptr, _ := readed_field.DataPtrUint8()
	delta := 0
	last_pixel_field := 0
	for a := 0; a < 480; a += 5 {
		last_pixel_field = 0
		//Vertical Verification of transations green-white-green
		for y := 0; y < 480; y++ {
			if image_field_ptr[y*480+a] == 0 {
				delta = y - last_pixel_field
				if MAX_LINE_WIDTH > delta && delta > MIN_LINE_WIDTH {
					x_dist := int(real_coordenates_omni[a][int(math.Round(float64(last_pixel_field+y)/2.0))][0]/10) + 80
					y_dist := -int(real_coordenates_omni[a][int(math.Round(float64(last_pixel_field+y)/2.0))][1]/10) + 80
					if 0 < x_dist && x_dist < 160 && 0 < y_dist && y_dist < 160 {
						readed_field_ptr[y_dist*160+x_dist] = 255
						last_pixel_field = y
					}else{
					last_pixel_field = y
					}
				} else {
					last_pixel_field = y
				}
			}
		}
		last_pixel_field = 0		
		//Horizontal Verification of transations green-white-green
		for x := 0; x < 480; x++ {
			if image_field_ptr[a*480+x] == 0 {
			
				delta = x - last_pixel_field
				if MAX_LINE_WIDTH > delta && delta > MIN_LINE_WIDTH {
					x_dist := int(real_coordenates_omni[last_pixel_field+int(delta/2)][a][0]/10) + 80
					y_dist := -int(real_coordenates_omni[last_pixel_field+int(delta/2)][a][1]/10) + 80
					if 0 < x_dist && x_dist < 160 && 0 < y_dist && y_dist < 160 {
						readed_field_ptr[y_dist*160+x_dist] = 255
						last_pixel_field = x
					}else{
					last_pixel_field = x
					}
				} else {
					last_pixel_field = x
				}
			}
		}
	}
}
