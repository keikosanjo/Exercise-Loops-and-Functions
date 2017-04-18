package main

import (
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}


func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{255, 0, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
