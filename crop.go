package imageutils

import (
	"errors"
	"image"
	"reflect"
)

// CropImage Crop an image using the specified cordinates
func CropImage(img image.Image, x0, x1, y1, y0 int) (image.Image, error) {
	if img == nil || reflect.ValueOf(img).Kind() == reflect.Ptr && reflect.ValueOf(img).IsNil() {
		return nil, errors.New("img is empty")
	}

	if x0 >= x1 || y0 >= y1 {
		return nil, errors.New("invalid dimension: not bounded by image")
	}
	newImgRGBA := image.NewNRGBA(image.Rect(x0, y0, x1, y1))
	croppedImg := newImgRGBA.SubImage(image.Rect(x0, y0, x1, y1)).(*image.NRGBA)
	*newImgRGBA = *croppedImg
	return newImgRGBA, nil
}
