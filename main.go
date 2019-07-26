package main

import (
	"log"

	"gocv.io/x/gocv"
)

func main() {
	video, err := gocv.OpenVideoCapture(0)
	if err != nil {
		log.Fatal(err)
	}

	img := gocv.NewMat()
	sobel := gocv.NewMat()
	w := gocv.NewWindow("normal")
	ws := gocv.NewWindow("sobel")

	for {
		video.Read(&img)
		gocv.Sobel(img, &sobel, 1, 1, 0, 3, 1, 0, gocv.BorderReflect101)

		w.IMShow(img)
		ws.IMShow(sobel)
		if w.WaitKey(1) == 27 || ws.WaitKey(1) == 27 {
			break
		}
	}
}
