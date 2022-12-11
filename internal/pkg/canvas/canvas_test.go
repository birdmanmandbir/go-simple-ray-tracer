package canvas

import (
	"github.com/birdmanmandbir/go-simple-ray-tracer/internal/pkg/mat"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	type args struct {
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want *Canvas
	}{
		{args: args{width: 10, height: 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCanvas(tt.args.width, tt.args.height)
			assert.Equal(t, tt.args.width, c.W)
			assert.Equal(t, tt.args.height, c.H)
			for i := 0; i < tt.args.width; i++ {
				for j := 0; j < tt.args.height; j++ {
					assert.Equal(t, mat.NewColor(0, 0, 0), c.ColorAt(i, j))
				}
			}
		})
	}
}

func TestCanvas_WritePixel(t *testing.T) {
	type fields struct {
		W      int
		H      int
		Pixels [][]*mat.Tuple4
	}
	type args struct {
		x     int
		y     int
		color *mat.Tuple4
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{fields: fields{W: 10, H: 20}, args: args{x: 2, y: 3, color: mat.NewColor(1, 0, 0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Canvas{
				W:      tt.fields.W,
				H:      tt.fields.H,
				Pixels: tt.fields.Pixels,
			}
			c.WritePixel(tt.args.x, tt.args.y, tt.args.color)
			assert.Equal(t, mat.NewColor(1, 0, 0), c.ColorAt(tt.args.x, tt.args.y))
		})
	}
}
