package canvas

import (
    "github.com/atmatm9182/wang/primitive"
)

type Canvas interface {
	DrawTriangle(primitive.Triangle)
}
