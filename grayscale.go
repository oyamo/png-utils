package imageutils

import (
	"errors"
	"image"
	"image/color"
	"log"
	"reflect"
	"sync"
)

func grayScaleImage(img image.Image, newImgRGBA *image.RGBA, routineCount int) {
	var workGroup sync.WaitGroup
	workGroup.Add(routineCount)

	for i := 0; i < routineCount; i++ {
		start := (newImgRGBA.Rect.Dx() / routineCount) * i
		end := ((newImgRGBA.Rect.Dx() / routineCount) * i) + (newImgRGBA.Rect.Dx() / routineCount)
		go func(startX, endX int) {
			for ix := startX; ix < endX; ix++ {
				for j := 0; j < newImgRGBA.Bounds().Dy(); j++ {
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

	newImgRGBA := image.NewRGBA(img.Bounds())
	grayScaleImage(img, newImgRGBA, 4)
	return newImgRGBA, nil
}
