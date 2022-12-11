package canvas

import (
	"github.com/birdmanmandbir/go-simple-ray-tracer/internal/pkg/mat"
	"github.com/stretchr/testify/assert"
	"strings"
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
	red := mat.NewColor(1, 0, 0)
	c := NewCanvas(10, 20)
	c.WritePixel(2, 3, red)
	assert.Equal(t, red, c.ColorAt(2, 3))
}

func TestCanvas_ToPPMHeader(t *testing.T) {
	c := NewCanvas(5, 3)
	ppm := strings.Join(strings.Split(c.ToPPM(), "\n")[0:3], "\n")
	expected := "P3\n5 3\n255"
	assert.Equal(t, expected, ppm)
}

func TestCanvas_ToPPM(t *testing.T) {
	c := NewCanvas(5, 3)
	color1 := mat.NewColor(1.5, 0, 0)
	color2 := mat.NewColor(0, 0.5, 0)
	color3 := mat.NewColor(-0.5, 0, 1)
	c.WritePixel(0, 0, color1)
	c.WritePixel(2, 1, color2)
	c.WritePixel(4, 2, color3)
	want :=
		`P3
5 3
255
255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`
	assert.Equal(t, want, c.ToPPM())
}
