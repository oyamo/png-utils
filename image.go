package imageutils

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

// ReadImage Reads an image file and returns a png.Image interface
func ReadImage(path string) (image.Image, error) {
	// read raw file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// read image
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

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
			rotation_x := float64(xn-centerx)*math.Cos(angleRadians) + float64(float32(yn)-centery)*math.Sin(angleRadians) + float64(centerx)
			rotation_y := float64(xn-centerx)*math.Sin(angleRadians) + float64(float32(yn)-centery)*math.Cos(angleRadians) + float64(centery)
			if rotation_x >= 0 && rotation_x < float64(bounds.Dx()) && rotation_y >= 0 && rotation_y < float64(bounds.Dy()) {
				colorAt := img.At(int(rotation_x), int(rotation_y))
				R, G, B, A := colorAt.RGBA()
				colorAtRGBA := color.RGBA{uint8(R), uint8(G), uint8(B), uint8(A)}
				newImg.SetRGBA(int(rotation_x), int(rotation_y), colorAtRGBA)
			} else {
				newImg.SetRGBA(int(rotation_x), int(rotation_y), color.RGBA{})
			}
		}
	}

	return newImg, nil
}
