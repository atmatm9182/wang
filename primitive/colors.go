package primitive

import (
    "math"
    "math/rand"
)

type Color [4]byte

var (
	red = Color{0xFF, 0x00, 0x00, 0xFF}
	white = Color{0xFF, 0xFF, 0xFF, 0xFF}
    blue = Color{0x00, 0x00, 0xFF, 0xFF}
    green = Color{0x00, 0xFF, 0x00, 0xFF}
    yellow = Color{0xFF, 0xFF, 0x00, 0xFF}
)

var Colors = map[string]Color{
    "red": red,
    "white": white,
    "blue": blue,
    "green": green,
    "yellow": yellow,
}

func (c Color) Bytes() [3]byte {
	res := [3]byte{}
	m := float64(c[3]) / 255.0
	for i := 0; i < 3; i++ {
		res[i] = byte(math.Round(float64(c[i]) * m))
	}
	return res
}

func RandColor() Color {
    var bs [4] byte
    for i := 0; i < 3; i++ {
        bs[i] = byte(rand.Intn(255))
    }
    bs[3] = byte(rand.Intn(127) + 128)
    return bs
}

