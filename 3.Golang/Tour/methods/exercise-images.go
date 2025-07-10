package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
}

func (self Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (self Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, self.w, self.h)
}

func (self Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{128, 128}
	pic.ShowImage(m)
}
