package opencv

import (
	"gocv.io/x/gocv"
)

func OpenCV001() {
	webcam, _ := gocv.OpenVideoCapture(0)
	window := gocv.NewWindow("hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
