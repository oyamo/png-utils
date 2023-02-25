package imageutils

import (
	"errors"
	"image"
	"image/color"
	"math"
)

func blend(watermark color.Color, main color.Color) color.Color {
	wr, wg, wb, wa := watermark.RGBA()
	mr, mg, mb, ma := main.RGBA()

	// If the watermark pixel is fully transparent, return the main pixel.
	if wa == 0 {
		return main
	}

	// If the watermark pixel is fully opaque, return the watermark pixel.
	if wa == 0xffff {
		return watermark
	}

	// Calculate the blended color using the alpha values of the two pixels.
	alpha := float64(wa) / float64(0xffff)
	r := uint16(float64(wr)*alpha + float64(mr)*(1-alpha))
	g := uint16(float64(wg)*alpha + float64(mg)*(1-alpha))
	b := uint16(float64(wb)*alpha + float64(mb)*(1-alpha))
	a := uint16(math.Max(float64(wa), float64(ma)))

	return color.RGBA64{R: r, G: g, B: b, A: a}
}

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
			waterMarkPixelColor := watermark.At(i-x, j-y)
			mainImagePixelColor := main.At(i, j)

			blendedColor := blend(waterMarkPixelColor, mainImagePixelColor)
			main.(*image.NRGBA).Set(i, j, blendedColor)
		}
	}

	return main, nil
}
