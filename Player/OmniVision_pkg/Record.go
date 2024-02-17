package OmniVision_pkg
/*
import (
	//"context"
	"fmt"
	//"image"
	//"image/color"
	//"log"
	//"math"
	//GK "player/GK_Position"
	"os"
	//coms "player/communication"
	//"runtime"
	//"time"
	"strconv"
	
	"gocv.io/x/gocv"
	//"google.golang.org/grpc"
	//"golang.org/x/tour/pic"
)

func (vision *OmniVision) Record2(x_enc int,y_enc int,x_vision int,y_vision int,x_final int ,y_final int) {
		
		//vision.pointer_camera = C.Get_Frame()
		

		
		//vision.record.Write(img)
		A := "A" + strconv.Itoa(vision.index_exc)
		B := "B" + strconv.Itoa(vision.index_exc)
		C := "C" + strconv.Itoa(vision.index_exc)
		D := "D" + strconv.Itoa(vision.index_exc)
		E := "E" + strconv.Itoa(vision.index_exc)
		F := "F" + strconv.Itoa(vision.index_exc)
		

		vision.file2.SetCellValue("Sheet1", A, x_enc)
		vision.file2.SetCellValue("Sheet1", B, y_enc)
		vision.file2.SetCellValue("Sheet1", C, x_vision)
		vision.file2.SetCellValue("Sheet1", D, y_vision)
		vision.file2.SetCellValue("Sheet1", E, x_final)
		vision.file2.SetCellValue("Sheet1", F, y_final)
		vision.index_exc2++
		

}
func (vision *OmniVision) ReadEX() ( int, int) {
		
		//vision.angle = coms.Get_bussola()
		//x, y := coms.GetDisplacement()
		//vision.pointer_camera = C.Get_Frame()
		//fmt.Println(x,y)

		
		//vision.record.Write(img)
		A := "A" + strconv.Itoa(vision.index_exc)
		B := "B" + strconv.Itoa(vision.index_exc)
		C := "C" + strconv.Itoa(vision.index_exc)

		//vision.file.SetCellValue("Sheet1", A, vision.angle)
		//vision.file.SetCellValue("Sheet1", B, x)
		//vision.file.SetCellValue("Sheet1", C, y)
		
		angle_,_ := vision.file.GetCellValue("Sheet1", A)
		x_, _ := vision.file.GetCellValue("Sheet1", B)
		y_, _ := vision.file.GetCellValue("Sheet1", C)
		angle, _ := strconv.Atoi(angle_)
		x, _ := strconv.Atoi(x_)
		y, _ := strconv.Atoi(y_)
		fmt.Println(angle,x,y)
		vision.index_exc++
		
		fmt.Println(vision.index_exc)
		key := gocv.WaitKey(50)
		fmt.Println("Erro6")
		if key == 27 {
			vision.StopRecord2()			
			os.Exit(0)
		}
	return  x, y
}


func (vision *OmniVision) StopRecord2(){
	//vision.record.Close()
	vision.file2.SetActiveSheet(vision.index_exc)
	vision.file2.SaveAs("DataLOC5.xlsx")
	os.Exit(0)}
	*/
