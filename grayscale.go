package imageutils

import (
	"errors"
	"image"
	"image/color"
	"reflect"
)

func GrayScaleImage(img image.Image) (image.Image, error) {
	if img == nil || reflect.ValueOf(img).Kind() == reflect.Ptr && reflect.ValueOf(img).IsNil() {
		return nil, errors.New("img is empty")
	}

	newImgRGBA := image.NewRGBA(img.Bounds())
	for i := 0; i < newImgRGBA.Bounds().Dx(); i++ {
		for j := 0; j < newImgRGBA.Bounds().Dy(); j++ {
			pixel := img.At(i, j)
			_, _, _, alpha := pixel.RGBA()
			newPixel := color.RGBA{A: uint8(alpha)}
			newImgRGBA.SetRGBA(i, j, newPixel)
		}
	}

	return newImgRGBA, nil
}
