package main

import (
	"github.com/atmatm9182/wang/canvas"
	"github.com/atmatm9182/wang/primitive"
	"github.com/atmatm9182/wang/tile"
)

func main() {
    c := canvas.NewPPMCanvas(1024, 1024)
    builder := tile.NewBuilder(&c)
    colors := [4]primitive.Color{
        primitive.Colors["red"],
        primitive.Colors["blue"],
        primitive.Colors["green"],
        primitive.Colors["yellow"],
    }
    builder.
        Colors(colors).
        Tiles(16).
        Build()
    c.Save("first_test.ppm")
}
