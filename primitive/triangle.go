package primitive

import (
    "slices"
)

type Triangle struct {
	Coords [3]Coord
	Color  Color
}

func sign(p1, p2, p3 Coord) float32 {
	return (p1.X()-p3.X())*(p2.Y()-p3.Y()) - (p2.X()-p3.X())*(p1.Y()-p3.Y())
}

func (t *Triangle) Contains(c Coord) bool {
	d1 := sign(c, t.Coords[0], t.Coords[1])
	d2 := sign(c, t.Coords[1], t.Coords[2])
	d3 := sign(c, t.Coords[2], t.Coords[0])

	has_neg := d1 < 0 || d2 < 0 || d3 < 0
	has_pos := d1 > 0 || d2 > 0 || d3 > 0

	return !(has_neg && has_pos)
}

func (t *Triangle) Bounds() (min_x, max_x, min_y, max_y float32) {
	xs := make([]float32, 0, 3)
	ys := make([]float32, 0, 3)

	for _, c := range t.Coords {
		xs = append(xs, c.X())
		ys = append(ys, c.Y())
	}

	min_x = slices.Min(xs)
	max_x = slices.Max(xs)
	min_y = slices.Min(ys)
	max_y = slices.Max(ys)

	return
}

