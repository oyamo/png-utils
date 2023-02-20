package imageutils

import (
	"image"
	"reflect"
)

func GrayScaleImage(img image.Image) (image.Image, error) {
	if img == nil || reflect.ValueOf(img).Kind() == reflect.Ptr && reflect.ValueOf(img).IsNil() {

	}
}
