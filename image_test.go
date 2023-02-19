package imageutils

import (
	"image"
	"testing"
)

func TestReadImage(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *image.Image
		wantErr bool
	}{
		{
			name: "Read simple image",
			args: args{
				path: "/home/oyamo/GolandProjects/image-utils/example/worship-experience.png",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ReadImage(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
