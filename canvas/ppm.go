package canvas

import (
	"fmt"
	"math"
	"os"

	. "github.com/atmatm9182/wang/primitive"
)

type PPMCanvas struct {
	Pixels [][]Color
	Width  uint
	Height uint
}

func (c *PPMCanvas) NormalizeCoord(coord Coord) (x, y uint) {
	fx := float64(coord.X())
	fy := float64(1.0 - coord.Y())

	fw := float64(c.Width)
	fh := float64(c.Height)
	x = uint(math.Round(fx * fw))
	y = uint(math.Round(fy * fh))

	return
}

func (c *PPMCanvas) NormalizeTriangle(t Triangle) Triangle {
	nt := Triangle{}
	nt.Color = t.Color
	for idx := range t.Coords {
		x, y := c.NormalizeCoord(t.Coords[idx])
		nt.Coords[idx] = Coord{float32(x), float32(y)}
	}
	return nt
}

func (c *PPMCanvas) DrawTriangle(t Triangle) {
	norm := c.NormalizeTriangle(t)
	min_x, max_x, min_y, max_y := norm.Bounds()
	for row := uint(min_y); row < uint(max_y); row++ {
		for col := uint(min_x); col < uint(max_x); col++ {
			cur_coord := Coord{float32(col), float32(row)}
			if norm.Contains(cur_coord) {
				c.Pixels[row][col] = t.Color
            }
		}
	}
}

func (c *PPMCanvas) Bytes() []byte {
	header := c.Header()
	res := make([]byte, 0, c.Width*c.Height*3+uint(len(header)))
	res = append(res, header...)
	for _, row := range c.Pixels {
		for _, pixel := range row {
			bs := pixel.Bytes()
			res = append(res, bs[:]...)
		}
	}
	return res
}

func (c *PPMCanvas) Header() []byte {
	return []byte(fmt.Sprintf("P6\n%d %d\n255\n", c.Width, c.Height))
}

func (c *PPMCanvas) Save(filename string) (err error) {
	var f *os.File
	f, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return
	}

	bytes := c.Bytes()

	var n int
	n, err = f.Write(bytes)

	if err != nil {
		return
	}

	if n != len(bytes) {
		err = fmt.Errorf("Writing to file %s returned %d bytes instead of expected %d", filename, n, len(bytes))
	}

	return
}

func NewPPMCanvas(w, h uint) PPMCanvas {
	pixels := make([][]Color, h)
	for i := range pixels {
		pixels[i] = make([]Color, w)
	}

	return PPMCanvas{
		Pixels: pixels,
		Width:  w,
		Height: h,
	}
}

