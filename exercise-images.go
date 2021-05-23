package main

import (
	"image/color"
	"image"
	"golang.org/x/tour/pic"
)

type Image struct {
	width, height int
	colors        uint8
}

func (image Image) ColorModel() color.Model {
	return color.RGBAModel
}


func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{ uint8(2)+uint8(x),uint8(200) + uint8(y), 1, 255}
}

func main() {
	m := Image{100, 100, 200}
	pic.ShowImage(m)
}
