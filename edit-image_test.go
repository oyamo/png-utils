package imageutils

import "testing"

func TestEditImage(t *testing.T) {
	type args struct {
		pathIn  string
		pathOut string
		op      Operation
		args    *OperationArgs
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Rotate image",
			args: args{
				op:      OperationRotate,
				args:    &OperationArgs{RotationDeg: 30},
				pathIn:  "/home/oyamo/GolandProjects/image-utils/example/worship-experience.png",
				pathOut: "/home/oyamo/GolandProjects/image-utils/example/worship-experience-rotate-30.png",
			},
			wantErr: false,
		}, {
			name: "Compress Image",
			args: args{
				op:      OperationLCompresstion,
				args:    &OperationArgs{RotationDeg: 30},
				pathIn:  "/home/oyamo/GolandProjects/image-utils/example/worship-experience.png",
				pathOut: "/home/oyamo/GolandProjects/image-utils/example/worship-experience-compress.png",
			},
			wantErr: false,
		}, {
			name: "GrayScale Image",
			args: args{
				op:      OperationGrayScale,
				args:    &OperationArgs{RotationDeg: 30},
				pathIn:  "/home/oyamo/GolandProjects/image-utils/example/worship-experience.png",
				pathOut: "/home/oyamo/GolandProjects/image-utils/example/worship-experience-gray.png",
			},
			wantErr: false,
		}, {
			name: "Crop Image",
			args: args{
				op:      OperationCrop,
				args:    &OperationArgs{Rect: struct{ X0, X1, Y0, Y1 int }{X0: 500, X1: 1500, Y0: 500, Y1: 1500}},
				pathIn:  "/home/oyamo/GolandProjects/image-utils/example/worship-experience.png",
				pathOut: "/home/oyamo/GolandProjects/image-utils/example/worship-cropped.png",
			},
			wantErr: false,
		}, {
			name: "Operation Label Image",
			args: args{
				op:      OperationAddLabel,
				args:    &OperationArgs{Rect: struct{ X0, X1, Y0, Y1 int }{X0: 500, X1: 1500, Y0: 500, Y1: 1500}},
				pathIn:  "/home/oyamo/GolandProjects/image-utils/example/enduro.png",
				pathOut: "/home/oyamo/GolandProjects/image-utils/example/enduro-labelled.png",
			},
			wantErr: false,
		}, {
			name: "Watermark Image",
			args: args{
				op:      OperationAddImageWatermark,
				args:    &OperationArgs{Rect: struct{ X0, X1, Y0, Y1 int }{X0: 500, X1: 1500, Y0: 500, Y1: 1500}},
				pathIn:  "/home/oyamo/GolandProjects/image-utils/example/enduro.png",
				pathOut: "/home/oyamo/GolandProjects/image-utils/example/enduro-watermarked.png",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := EditImage(tt.args.pathIn, tt.args.pathOut, tt.args.op, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("EditImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func BenchEditImage(t *testing.B) {
	type args struct {
		pathIn  string
		pathOut string
		op      Operation
		args    *OperationArgs
	}

	arg := args{
		op:      OperationRotate,
		args:    &OperationArgs{RotationDeg: 30},
		pathIn:  "/home/oyamo/GolandProjects/image-utils/example/worship-experience.png",
		pathOut: "/home/oyamo/GolandProjects/image-utils/example/worship-experience-rotate-30.png",
	}

	err := EditImage(arg.pathIn, arg.pathOut, OperationGrayScale, arg.args)
	if err != nil {
		return
	}
}
