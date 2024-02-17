package OmniVision_pkg

import (
	"image"
	"image/color"
	"gocv.io/x/gocv"
)


// Constant to tranform Pixels in our Range X
const PIXEL_RANGE_X float64 = 1000 / ((A / 2) / 10)

// Constant to tranform Pixels in our Range Y
const PIXEL_RANGE_Y float64 = 1000 / ((B / 2) / 10)

// Pixels_to_MM transform pixels position in millimeters position. 
// In other words, it transforms the position in the field image into distances in the real world
// 
// Parameters:
//
//   - `x` : int -> X in Pixels
//   - `y` : int -> Y in Pixels
//
// Returns:
//
//   - `x` : int -> X in mm
//   - `y` : int -> Y in mm
//
func Pixels_to_MM(x, y int) (int, int) {
	return int(100 * -(y - 310)), int(100 * -(x - 270))
}

// CM_to_Pixels tranform centimetres position in pixels position.
// In other words, the opposite of the Pixels_to_MM function.
// 
// Parameters:
//
//   - `x` : float64 -> X in cm
//   - `y` : float64 -> Y in cm
//
// Returns:
//
//   - `x` : int -> X in Pixels
//   - `y` : int -> Y in Pixels
//
func CM_to_Pixels(x, y float64) (int, int) {
	return int(-(y / 10) + 270), int(-(x / 10) + 310)
}

// The ArgMax function gets the index of the largest element in an array.
//
// Parameters:
//
//   - `v` : []int -> Array
//
// Returns:
//
//   - `max_index` : int -> Index of max value
//
func ArgMax(v []int) int {
	m := 0
	max_index := 0
	for i, e := range v {
		if i == 0 || e > m {
			m = e
			max_index = i
		}
	}
	return max_index
}

// Rotate is a function that rotates the image to a specific angle using nearest neighbour interpolation. 
//
// Parameters:
//
//   - `original` : gocv.Mat -> Matrix to rotate
//   - `rotated` : *gocv.Mat -> Rotated Matrix
//   - `angle` : int -> rotation angle
//
func Rotate(original gocv.Mat, rotated *gocv.Mat, angle int) {
	cols, rows := original.Cols(), original.Rows()
	center := image.Point{cols / 2, rows / 2}
	rotation := gocv.GetRotationMatrix2D(center, float64(-angle), 1.0)
	gocv.WarpAffineWithParams(original, rotated, rotation, image.Point{cols, rows}, gocv.InterpolationNearestNeighbor, gocv.BorderConstant, color.RGBA{0, 0, 0, 0})
}
