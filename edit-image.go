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
	OperationCrop
	OperationAddLabel
	OperationAddImageWatermark
)

type OperationArgs struct {
	RotationDeg int
	Rect        struct {
		X0, X1, Y0, Y1 int
	}
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
	case OperationCrop:
		newImage, err = CropImage(img, args.Rect.X0, args.Rect.X1, args.Rect.Y0, args.Rect.Y1)
		break
	case OperationAddLabel:
		newImage, err = AddLabel(img, 24, 24, "Dream", 0xFF20FF00)
		break
	case OperationAddImageWatermark:
		watermarkImage, err := ReadImage("/home/oyamo/GolandProjects/image-utils/example/golang_logo_icon_171073.png")
		if err != nil {
			return err
		}
		newImage, err = AddWatermarkImage(img, watermarkImage, 128, 128)

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
