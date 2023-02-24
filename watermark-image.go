package imageutils

import (
	"errors"
	"image"
	"image/color"
)

func AddWatermarkImage(main image.Image, watermark image.Image, x, y int) (image.Image, error) {
	mainImageHeight := main.Bounds().Dy()
	mainImageWidth := main.Bounds().Dx()

	watermarkImageHeight := watermark.Bounds().Dy()
	watermarkImageWidth := watermark.Bounds().Dx()

	if x < 0 || y < 0 {
		return nil, errors.New("dimensions out of bounds")
	}

	if x > mainImageWidth || y > mainImageHeight {
		return nil, errors.New("dimensions out of bounds")
	}

	for i := x; i < watermarkImageWidth+x; i++ {
		for j := y; j < watermarkImageHeight+y; j++ {
			waterMarkPixel := watermark.At(i-x, j-y)
			mainImagePixel := main.At(i, j)
			r, g, b, opacity := waterMarkPixel.RGBA()
			waterMarkColor := color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(opacity)}
			r, g, b, o2 := mainImagePixel.RGBA()
			mainImageColor := color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(o2)}
			if opacity > 20 {
				main.(*image.NRGBA).Set(i, j, waterMarkColor)
			} else if opacity > 5 {
				main.(*image.NRGBA).Set(i, j, mainImageColor)
			} else {
				main.(*image.NRGBA).Set(i, j, mainImageColor)
			}
		}
	}

	return main, nil
}
