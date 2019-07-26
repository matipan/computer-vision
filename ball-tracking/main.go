package main

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

var (
	rcolor = color.RGBA{G: 255, A: 255}
	lcolor = color.RGBA{R: 255, A: 255}

	wt   = gocv.NewWindow("thersholded")
	wi   = gocv.NewWindow("images")
	img  = gocv.NewMat()
	mask = gocv.NewMat()
)

func main() {
	lhsv := gocv.Scalar{Val1: 49, Val2: 89, Val3: 0}
	hhsv := gocv.Scalar{Val1: 109, Val2: 255, Val3: 255}

	video, _ := gocv.OpenVideoCapture(0)
	wt.ResizeWindow(600, 600)
	wt.MoveWindow(0, 0)
	wi.MoveWindow(600, 0)
	wi.ResizeWindow(600, 600)
	frame := gocv.NewMat()
	hsv := gocv.NewMat()
	kernel := gocv.NewMat()

	queue := New(40)

	for {
		video.Read(&img)
		gocv.Flip(img, &img, 1)
		gocv.Resize(img, &img, image.Point{X: 600, Y: 600}, 0, 0, gocv.InterpolationLinear)
		gocv.GaussianBlur(img, &frame, image.Point{X: 11, Y: 11}, 0, 0, gocv.BorderReflect101)
		gocv.CvtColor(frame, &hsv, gocv.ColorBGRToHSV)

		gocv.InRangeWithScalar(hsv, lhsv, hhsv, &mask)
		gocv.Erode(mask, &mask, kernel)
		gocv.Dilate(mask, &mask, kernel)

		cnt := bestContour(mask, 2000)
		if len(cnt) == 0 {
			queue.Clear()
			if imShow() {
				break
			}
			continue
		}
		rect := gocv.BoundingRect(cnt)
		x, y := middle(rect)
		queue.Push(image.Point{X: x, Y: y})

		queue.RangeWithPrevious(func(c image.Point, p image.Point) {
			gocv.Line(&img, p, c, lcolor, 2)
		})

		gocv.Rectangle(&img, rect, rcolor, 2)
		if imShow() {
			break
		}
	}

}

func imShow() bool {
	wi.IMShow(img)
	wt.IMShow(mask)
	return wi.WaitKey(1) == 27 || wt.WaitKey(1) == 27
}

// bestContour obtains the biggest contour in the frame(provided is bigger)
// than the minArea.
func bestContour(frame gocv.Mat, minArea float64) []image.Point {
	cnts := gocv.FindContours(frame, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	var (
		bestCnt  []image.Point
		bestArea = minArea
	)
	for _, cnt := range cnts {
		if area := gocv.ContourArea(cnt); area > bestArea {
			bestArea = area
			bestCnt = cnt
		}
	}
	return bestCnt
}

// middle calculates the middle x and y of a rectangle.
func middle(rect image.Rectangle) (x int, y int) {
	return (rect.Max.X-rect.Min.X)/2 + rect.Min.X, (rect.Max.Y-rect.Min.Y)/2 + rect.Min.Y
}
