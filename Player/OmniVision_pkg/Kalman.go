package OmniVision_pkg
import (
	"math"
	coms "player/communication"
	"gonum.org/v1/gonum/mat"
)

// Contain all variables needed in Kalman filter.
type Kalman struct {
	X      *mat.VecDense // State
	P      *mat.Dense    // Covariance
	K      *mat.Dense    // Gain
	Q      *mat.Dense    // Process Noise Covariance
	R      *mat.Dense    // Measurement Covariance
}

// Init_Kalman is a function that init all variables of Kalman filter.
// It is a method of the OmniVision struct.
//
// Parameters:
//
//   - `x` : float64 ->  Initial position in X
//   - `y` : float64 ->  Initial position in Y
//
func (vision *OmniVision) Init_Kalman(x, y float64) {
	var new_fk Kalman
	vision.kalman = &new_fk

	vision.kalman.Q = mat.NewDense(2, 2, []float64{
		0.5, 0,
		0, 0.5,
	})
	vision.kalman.R = mat.NewDense(2, 2, []float64{
		3000, 0,
		0, 3000,
	})
	
	vision.kalman.X = mat.NewVecDense(2, []float64{
		x, y,
	})
	vision.kalman.K = mat.NewDense(2, 2, []float64{
		0, 0,
		0, 0,
	})
	const initp float64 = 1
	vision.kalman.P = mat.NewDense(2, 2, []float64{
		initp, 0,
		0, initp,
	})
	

}

// Reset_Kalman is a function that reset the position of Kalman filter.
// It is a method of the OmniVision struct.
//
func (vision *OmniVision) Reset_Kalman() {
	
	vision.kalman.X.SetVec(0, vision.x)
	vision.kalman.X.SetVec(1, vision.y)
}

// Update_Kalman is a function that update Kalman filter using Computer Vision and Encoders values.
// It is a method of the OmniVision struct.
//
// Returns:
//
//   - `x` : float64 ->  Final position in X
//   - `y` : float64 ->  Final position in Y
//
func (vision *OmniVision) Update_Kalman() (float64, float64) {
	
	x_enc_rel, y_enc_rel :=  coms.GetDisplacement()// vision.ReadEX()
	to_RAD := 0.017453293
	rad_y_angle := float64(vision.angle) * to_RAD
	rad_x_angle := float64(vision.angle-90) * to_RAD

	x_enc := float64(y_enc_rel)*math.Cos(rad_y_angle) + float64(x_enc_rel)*math.Cos(rad_x_angle)
	y_enc := float64(y_enc_rel)*math.Sin(rad_y_angle) + float64(x_enc_rel)*math.Sin(rad_x_angle)
	
	//Prediction:	
	//Xpriori = A*lastX + B*U, A=I and B=I => Xpriori = lastX + encoders
	vision.kalman.X = mat.NewVecDense(2, []float64{ vision.kalman.X.AtVec(0) + x_enc, vision.kalman.X.AtVec(1) - y_enc,})
	vision_Mat := mat.NewVecDense(2, []float64{float64(vision.x_vision), float64(vision.y_vision),})
	vision.x_enc=int(x_enc)
	vision.y_enc=int(y_enc)
	
	//Ppriori = A*lastP*A' + Q => P = P + Q
	vision.kalman.P.Add(vision.kalman.P, vision.kalman.Q)
	
	//Correction:
	var P_R mat.Dense
	var invP_R mat.Dense
	var K_V_X mat.VecDense
	var K_P mat.Dense
	
	//K = Ppriori*H'/(H*Ppriori*H'+R); => K = P/P+R
	P_R.Add(vision.kalman.P, vision.kalman.R)
	invP_R.Inverse(&P_R)
	vision.kalman.K.Mul(vision.kalman.P, &invP_R)
	
	//X = Xpriori + K*([visionX;visionY]-H*Xpriori); => X = X + K(vision-X)
	var V_X mat.VecDense
	V_X.SubVec(vision_Mat, vision.kalman.X)
	K_V_X.MulVec(vision.kalman.K, &V_X)
	vision.kalman.X.AddVec(vision.kalman.X, &K_V_X)
	
	//  P = P - K*H*P; => P = P - K*P
	K_P.Mul(vision.kalman.K, vision.kalman.P)
	vision.kalman.P.Sub(vision.kalman.P, &K_P)
	
	// mm to cm
	vision.x = vision.kalman.X.AtVec(0) / 10.0
	vision.y = vision.kalman.X.AtVec(1) / 10.0
	return vision.x, vision.y
}
