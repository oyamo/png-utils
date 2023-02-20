package imageutils

import (
	"errors"
	"image"
)

const CompressionPercentage = 80

func LossLessCompress(img image.Image) (image.Image, error) {
	if img == nil {
		return nil, errors.New("image is nil")
	}
	bounds := img.Bounds()
	newBounds := image.Rect(0, 0, bounds.Dx()*(CompressionPercentage/100), bounds.Dy()*(CompressionPercentage/100))
	newImage := image.NewRGBA(newBounds)
	for i := 0; i < newBounds.Dx(); i++ {
		for j := 0; j < newBounds.Dy(); j++ {
			atX := float32(i) * (1 + (CompressionPercentage / 100))
			atY := float32(j) * (1 + (CompressionPercentage / 100))
			pixelAt := img.At(int(atX), int(atY))
			newImage.Set(int(atY), int(atY), pixelAt)
		}
	}

	return newImage, nil
}
