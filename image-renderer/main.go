package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

// ColorModel returns the Image's color model.
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (i Image) Bounds() image.Rectangle {
	return image.Rectangle{Min: image.Point{0, 0},
		Max: image.Point{100, 100}}
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (i Image) At(x, y int) color.Color {
	//var factorX uint8 = uint8(x)
	//var factorY uint8 = uint8(y)
	return i.ColorModel().
		Convert(color.RGBA{uint8(0xFF&x) * 2,
			uint8(0xFF & x) + 50,
			uint8(0xFF & y),
			0xFF})

}
func main() {
	m := Image{}
	pic.ShowImage(m)
}
