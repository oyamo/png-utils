package imageutils

import (
	"errors"
	"image"
	"image/color"
	"reflect"
)

const CompressionPercentage float32 = 10.0

func LossLessCompress(img image.Image) (image.Image, error) {
	if img == nil || reflect.ValueOf(img).Kind() == reflect.Ptr && reflect.ValueOf(img).IsNil() {
		return nil, errors.New("image is nil")
	}
	bounds := img.Bounds()
	newBounds := image.Rect(0, 0, int(float32(bounds.Dx())*(CompressionPercentage/100.0)), int(float32(bounds.Dy())*(CompressionPercentage/100.0)))
	newImage := image.NewRGBA(newBounds)
	for i := 0; i < newBounds.Dx(); i++ {
		for j := 0; j < newBounds.Dy(); j++ {
			atX := int(float64(i) * float64(bounds.Dx()) / float64(newBounds.Dx()))
			atY := int(float64(j) * float64(bounds.Dy()) / float64(newBounds.Dy()))
			colorAt := img.At(int(atX), int(atY))
			R, G, B, A := colorAt.RGBA()
			colorAtRGBA := color.RGBA{R: uint8(R), G: uint8(G), B: uint8(B), A: uint8(A)}
			newImage.SetRGBA(int(i), int(j), colorAtRGBA)
		}
	}

	return newImage, nil
}
