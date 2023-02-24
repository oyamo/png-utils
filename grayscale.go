package imageutils

import (
	"errors"
	"image"
	"image/color"
	"log"
	"reflect"
	"sync"
)

// cropImage contains workers for grayscaling an image
func grayScaleImage(img image.Image, newImgRGBA *image.NRGBA, routineCount int) {
	var workGroup sync.WaitGroup
	workGroup.Add(routineCount)
	height := newImgRGBA.Rect.Dy()
	width := newImgRGBA.Rect.Dx()
	for i := 0; i < routineCount; i++ {
		start := (width / routineCount) * i
		end := ((width / routineCount) * i) + (width / routineCount)
		go func(startX, endX int) {
			for ix := startX; ix < endX; ix++ {
				for j := 0; j < height; j++ {
					pixel := img.At(ix, j)
					r, g, b, a := pixel.RGBA()
					// convert to grayscale using the formula: gray = 0.2126*r + 0.7152*g + 0.0722*b
					gray := uint8((19595*r + 38469*g + 7472*b + 1<<15) >> 24)
					newPixel := color.RGBA{R: gray, G: gray, B: gray, A: uint8(a)}
					newImgRGBA.Set(ix, j, newPixel)
				}
			}
			workGroup.Done()
			log.Println("DONE", startX, endX)

		}(start, end)
	}
	workGroup.Wait()
}

// GrayScaleImage Generates grayscale version of an image
// https://en.wikipedia.org/wiki/Grayscale
func GrayScaleImage(img image.Image) (image.Image, error) {
	if img == nil || reflect.ValueOf(img).Kind() == reflect.Ptr && reflect.ValueOf(img).IsNil() {
		return nil, errors.New("img is empty")
	}

	newImgRGBA := img.(*image.NRGBA)
	grayScaleImage(img, newImgRGBA, 12)
	return newImgRGBA, nil
}
