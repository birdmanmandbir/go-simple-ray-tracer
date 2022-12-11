package canvas

import (
	"fmt"
	"github.com/birdmanmandbir/go-simple-ray-tracer/internal/pkg/mat"
	"math"
	"strconv"
	"strings"
)

const MaxColor = 255.0
const MaxLineWidth = 69

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

func (c *Canvas) ToPPM() string {
	header := fmt.Sprintf("P3\n%d %d\n255\n", c.W, c.H)
	result := header
	colNumber := 0
	for i := 0; i < c.W; i++ {
		buf := strings.Builder{}
		for j := 0; j < c.H; j++ {
			clamped := clamp(c.ColorAt(i, j))
			writeColorToBuffer(clamped, &buf, &colNumber)
		}
		out := buf.String()
		out = strings.TrimSuffix(out, " ")
		result += out + "\n"
	}
	result = strings.TrimSuffix(result, "\n")
	return result
}

func clamp(color *mat.Tuple4) *mat.Tuple4 {
	newColor := mat.NewColor(0, 0, 0)
	for i := 0; i < 3; i++ {
		c := color.Get(i) * MaxColor
		rounded := math.Round(c)
		if rounded < 0 {
			rounded = 0
		} else if rounded > MaxColor {
			rounded = MaxColor
		}
		newColor.Data[i] = rounded
	}
	return newColor
}

func writeColorToBuffer(color *mat.Tuple4, buf *strings.Builder, colNumber *int) {
	for i := 0; i < 3; i++ {
		// Every color need 3 space(max), for example 255
		if *colNumber+3 > MaxLineWidth {
			buf.WriteString("\n")
			buf.WriteString(strconv.Itoa(int(color.Get(i))))
			buf.WriteString(" ")
			*colNumber = 3
		} else {
			buf.WriteString(strconv.Itoa(int(color.Get(i))))
			if *colNumber+6 < MaxLineWidth {
				buf.WriteString(" ")
			}
			*colNumber += 4
		}
	}
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
