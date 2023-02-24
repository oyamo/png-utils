package imageutils

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/inconsolata"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
)

// Add a a text
func AddLabel(img image.Image, x, y int, label string, labelColor uint32) (image.Image, error) {
	col := color.RGBA{
		A: uint8((labelColor >> 24) & 0xFF),
		R: uint8((labelColor >> 16) & 0xFF),
		G: uint8((labelColor >> 8) & 0xFF),
		B: uint8(labelColor & 0xFF),
	}

	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}
	d := font.Drawer{
		Dot:  point,
		Dst:  img.(*image.NRGBA),
		Face: inconsolata.Bold8x16,
		Src:  image.NewUniform(col),
	}

	d.DrawString(label)

	return img, nil
}
