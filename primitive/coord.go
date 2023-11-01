package primitive

type Coord [2]float32

func (c Coord) X() float32 {
	return c[0]
}

func (c Coord) Y() float32 {
	return c[1]
}

