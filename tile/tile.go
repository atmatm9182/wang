package tile

import (
    . "github.com/atmatm9182/wang/primitive"
)

type WangTile struct {
    Colors [4]Color
}

func (wt WangTile) Top() Color {
    return wt.Colors[0]
}

func (wt WangTile) Right() Color {
    return wt.Colors[1]
}

func (wt WangTile) Bot() Color {
    return wt.Colors[2]
}

func (wt WangTile) Left() Color {
    return wt.Colors[3]
}
