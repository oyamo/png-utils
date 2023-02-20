package imageutils

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

// RotateImage Rotates the image deg degrees about the center of the image
func RotateImage(img image.Image, deg int) (image.Image, error) {
	if img == nil {
		return nil, fmt.Errorf("img is nil")
	}
	bounds := img.Bounds()
	// convert angle to radians
	angleRadians := float64(float32(deg) * math.Pi / 180.0)
	// determine center of the image
	centerx, centery := float32(bounds.Dx()/2), float32(bounds.Dy()/2)
	// Create new image with bounds
	newImg := image.NewRGBA(img.Bounds())

	for xn := float32(0); xn < float32(bounds.Dx()); xn++ {
		for yn := 0; yn < bounds.Dy(); yn++ {
			//Apply the inverse rotation matrix to get the corresponding pixel in the original image
			rotationX := float64(xn-centerx)*math.Cos(angleRadians) + float64(float32(yn)-centery)*math.Sin(angleRadians) + float64(centerx)
			rotationY := float64(xn-centerx)*math.Sin(angleRadians) + float64(float32(yn)-centery)*math.Cos(angleRadians) + float64(centery)
			if rotationX >= 0 && rotationX < float64(bounds.Dx()) && rotationY >= 0 && rotationY < float64(bounds.Dy()) {
				colorAt := img.At(int(rotationX), int(rotationY))
				R, G, B, A := colorAt.RGBA()
				colorAtRGBA := color.RGBA{uint8(R), uint8(G), uint8(B), uint8(A)}
				newImg.SetRGBA(int(rotationX), int(rotationY), colorAtRGBA)
			} else {
				newImg.SetRGBA(int(rotationX), int(rotationY), color.RGBA{})
			}
		}
	}

	return newImg, nil
}
