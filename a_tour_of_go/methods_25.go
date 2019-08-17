package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type myImage struct{}

func (i myImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

func (i myImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (i myImage) At(x, y int) color.Color {
	sum := uint8(x ^ y)
	return color.RGBA{sum, sum, 255, 255}
}

// ShowImage (methods 25)
func ShowImage() {
	m := myImage{}
	pic.ShowImage(m)
}
