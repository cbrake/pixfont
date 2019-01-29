package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	txt := "ai"
	width := Font.MeasureString(txt)
	fmt.Println("MeasureString returned: ", width)
	//height := tightpixel15.Font.GetHeight()
	height := 10
	fmt.Println("width: ", width)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	widthActual := Font.DrawString(img, 0, 0, txt, color.Black)
	fmt.Println("DrawString returned: ", widthActual)
	imgS := img.SubImage(image.Rectangle{
		image.Point{0, 0},
		image.Point{widthActual, height},
	})
	fmt.Println("widthActual: ", widthActual)

	f, _ := os.OpenFile("hello.png", os.O_CREATE|os.O_RDWR, 0644)
	png.Encode(f, imgS)
}
