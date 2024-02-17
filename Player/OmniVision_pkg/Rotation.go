package OmniVision_pkg
import "gocv.io/x/gocv"

// The adjustAngle function chooses the best orientation for the image using the field lines as a reference. 
//
// Parameters:
//
//   - `rotated` : *gocv.Mat -> Rotated Matrix
//
// Outputs:
//
//   - `angle` : int -> Robot calculated angle
//
func (vision *OmniVision)adjustAngle(image_r *gocv.Mat) int {
	

	MIN_HISTOGRAM := -10
	MAX_HISTOGRAM := 10
	rotated := gocv.Zeros(160,160,gocv.MatTypeCV8UC1)
	defer rotated.Close()
	var max_values []int

	Rotate(*image_r, image_r, vision.angle)
	histH := gocv.Zeros(1,160,gocv.MatTypeCV32SC1)
	histV := gocv.Zeros(160,1,gocv.MatTypeCV32SC1)
	
	// # Clean Outside Lines
	vision.Clean_OutField(image_r, int(vision.x), int(vision.y))
	
	for i := MIN_HISTOGRAM; i <= MAX_HISTOGRAM; i++ {
		Rotate(*image_r, &rotated, i)
		gocv.Reduce(rotated, &histV, 0, gocv.ReduceSum, gocv.MatTypeCV32SC1)
		gocv.Reduce(rotated, &histH, 1, gocv.ReduceSum, gocv.MatTypeCV32SC1)
		_,maxValV, _, _ :=gocv.MinMaxLoc(histV)
		_,maxValH, _, _ :=gocv.MinMaxLoc(histH)
		max_values = append(max_values, int(maxValV+maxValH))
	}
	
	adjust := (ArgMax(max_values))+MIN_HISTOGRAM
	vision.angle+=adjust
	if -180 > vision.angle {
		vision.angle += 360
	}
	if vision.angle > 180 {
		vision.angle -= 360
	}
	
	Rotate(*image_r, image_r, adjust)
	
	return vision.angle
}
