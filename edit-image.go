package imageutils

import (
	"errors"
	"image"
	"reflect"
)

const (
	OperationLCompresstion = Operation(iota)
	OperationRotate
	OperationGrayScale
)

type OperationArgs struct {
	RotationDeg int
}

type Operation int

func EditImage(pathIn, pathOut string, op Operation, args *OperationArgs) error {
	var newImage image.Image

	if args == nil {
		return errors.New("missing operation arguments")
	}

	img, err := ReadImage(pathIn)
	if err != nil {
		return err
	}

	switch op {
	case OperationLCompresstion:
		newImage, err = LossLessCompress(img)
		break
	case OperationRotate:
		newImage, err = RotateImage(img, args.RotationDeg)
		break
	case OperationGrayScale:
		newImage, err = GrayScaleImage(img)
		break
	}

	if err != nil {
		return err
	}

	if newImage == nil || reflect.ValueOf(newImage).Kind() == reflect.Ptr &&
		reflect.ValueOf(newImage).IsNil() {
		return errors.New("fatal error")
	}

	err = SaveImage(newImage, pathOut)
	if err != nil {
		return err
	}

	return nil
}
