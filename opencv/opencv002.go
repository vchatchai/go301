package opencv

import (
	"fmt"
	"image/color"

	"gocv.io/x/gocv"
)

func OpenCV002() {
	deviceID := 0

	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	blue := color.RGBA{0, 0, 255, 0}

	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("haarcascade_frontalface_default.xml") {
		fmt.Println("Error reading  cascade file: haarcascade_frontalface_default.xml")
		return
	}

	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d faces\n", len(rects))

		for _, r := range rects {
			gocv.Rectangle(&img, r, blue, 3)
		}

		window.IMShow(img)
		window.WaitKey(1)
	}
}
