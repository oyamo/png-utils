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

func SaveImage(img image.Image, path string) error {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}
	encoder := png.Encoder{CompressionLevel: png.BestSpeed}
	err = encoder.Encode(file, img)
	if err != nil {
		return err
	}
	return nil
}
