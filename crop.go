package imageutils

import (
	"errors"
	"image"
	"log"
	"reflect"
	"sync"
)

// cropImage crop an image using workers
func cropImage(img image.Image, newImgRGBA *image.NRGBA, x0, x1, y0, y1 int, routineCount int) {
	var workGroup sync.WaitGroup
	workGroup.Add(routineCount)

	for i := 0; i < routineCount; i++ {
		start := (newImgRGBA.Rect.Dx() / routineCount) * i
		end := ((newImgRGBA.Rect.Dx() / routineCount) * i) + (newImgRGBA.Rect.Dx() / routineCount)
		go func(startX, endX int) {
			for ix := startX; ix < endX; ix++ {
				for j := 0; j < newImgRGBA.Bounds().Dy(); j++ {
					if ix >= x0 && ix < x1 && j >= y0 && j < y1 {
						pixel := img.At(ix, j)
						newImgRGBA.Set(ix, j, pixel)
					}
				}
			}
			workGroup.Done()
			log.Println("DONE", startX, endX)

		}(start, end)
	}
	workGroup.Wait()

	croppedImg := newImgRGBA.SubImage(image.Rect(x0, y0, x1, y1)).(*image.NRGBA)
	*newImgRGBA = *croppedImg
}

// CropImage Crop an image using the specified cordinates
func CropImage(img image.Image, x0, x1, y0, y1 int) (image.Image, error) {
	if img == nil || reflect.ValueOf(img).Kind() == reflect.Ptr && reflect.ValueOf(img).IsNil() {
		return nil, errors.New("img is empty")
	}

	if x0 >= x1 || y0 >= y1 {
		return nil, errors.New("invalid dimension: not bounded by image")
	}
	newImgRGBA := image.NewNRGBA(image.Rect(x0, y0, x1, y1))
	cropImage(img, newImgRGBA, x0, x1, y0, y1, 4)
	return newImgRGBA, nil
}
