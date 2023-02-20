package imageutils

import (
	"image"
	"image/png"
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
