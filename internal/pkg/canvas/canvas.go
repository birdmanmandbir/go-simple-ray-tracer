package canvas

import "github.com/birdmanmandbir/go-simple-ray-tracer/internal/pkg/mat"

type Canvas struct {
	W      int
	H      int
	Pixels [][]*mat.Tuple4
}

func (c *Canvas) ColorAt(x, y int) *mat.Tuple4 {
	return c.Pixels[x][y]
}

func (c *Canvas) checkWritePixel(x int, y int) bool {
	if x < 0 || y < 0 || x >= c.W || y >= c.H {
		println("WritePixel: Out of bound")
		return false
	}
	return true
}

func (c *Canvas) WritePixel(x int, y int, color *mat.Tuple4) {
	if !c.checkWritePixel(x, y) {
		return
	}
	c.Pixels[x][y] = color
}

func NewCanvas(width, height int) *Canvas {
	pixels := make([][]*mat.Tuple4, width)
	for i := 0; i < width; i++ {
		pixels[i] = make([]*mat.Tuple4, height)
		for j := 0; j < height; j++ {
			pixels[i][j] = mat.NewColor(0, 0, 0)
		}
	}
	return &Canvas{W: width, H: height, Pixels: pixels}
}
