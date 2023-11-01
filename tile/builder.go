package tile

import (
	"math"
	"math/rand"
	"sync"

	"github.com/atmatm9182/wang/canvas"
	. "github.com/atmatm9182/wang/primitive"
)

type Builder struct {
    canvas canvas.Canvas
    colors [4]Color
    tiles []WangTile
}

func at(w WangTile, coords [4]Coord) [4]Triangle {
    center := Coord{
        (coords[1].X() - coords[0].X()) / 2 + coords[0].X(),
        (coords[0].Y() - coords[3].Y()) / 2 + coords[2].Y(),
    }

    t := Triangle{
        Coords: [3]Coord{
            coords[0],
            center,
            coords[1],
        }, 
        Color: w.Colors[0],
    }

    r := Triangle{
        Coords: [3]Coord{
            coords[1],
            center,
            coords[2],
        }, 
        Color: w.Colors[1],
    }
    b := Triangle{
        Coords: [3]Coord{
            coords[2],
            center,
            coords[3],
        }, 
        Color: w.Colors[2],
    }
    l := Triangle{
        Coords: [3]Coord{
            coords[3],
            center,
            coords[0],
        }, 
        Color: w.Colors[3],
    }

    return [4]Triangle{t, r, b, l}
}

func (b *Builder) randTile() WangTile {
    w := WangTile{}
    w.Colors = [4]Color(b.randColors(4))
    return w
}

func (b *Builder) Colors(cs [4]Color) *Builder {
    b.colors = cs
    return b
}

func (b *Builder) randColors(n uint) []Color {
    res := make([]Color, n)
    for i := range res {
        ci := rand.Intn(4)
        res[i] = b.colors[ci]
    }

    return res
} 

func (b *Builder) genRight(leftColor Color) WangTile {
    wt := WangTile{}
    wt.Colors[3] = leftColor

    gcs := b.randColors(3)
    for i := 0; i < 3; i++ {
        wt.Colors[i] = gcs[i]
    }

    return wt
}

func (b *Builder) genBot(topColor Color) WangTile {
    wt := WangTile{}
    wt.Colors[0] = topColor

    gcs := b.randColors(3)
    for i := 1; i < 4; i++ {
        wt.Colors[i] = gcs[i - 1]
    }

    return wt
}

func (b *Builder) firstRow(n uint) []WangTile {
    res := make([]WangTile, 0, n)
    cur := b.randTile()
    res = append(res, cur)

    for i := uint(1); i < n; i++ {
        g := b.genAdjacentTile(cur, right)
        res = append(res, g)
        cur = g
    }

    return res
}

func (b *Builder) genAdjacentTile(wt WangTile, dir direction) WangTile {
    switch dir {
    case bot:
        return b.genBot(wt.Bot())
    case right:
        return b.genRight(wt.Right())
}
    panic("unreachable")
}

func (b *Builder) genBotRight(topColor, leftColor Color) WangTile {
    wt := WangTile{}
    wt.Colors[0] = topColor
    wt.Colors[3] = leftColor

    rc := b.randColors(2)
    for i := 1; i < 3; i++ {
        wt.Colors[i] = rc[i - 1]
    }

    return wt
}

type direction uint8
const (
    bot direction = iota
    right
)

func (b *Builder) nextRow(row []WangTile) []WangTile {
    res := make([]WangTile, 0, len(row))
    cur := b.genAdjacentTile(row[0], bot)
    res = append(res, cur)

    for i := 1; i < len(row); i++ {
        g := b.genBotRight(row[i].Bot(), cur.Right())
        res = append(res, g)
        cur = g
    }

    return res
}

func (b *Builder) genTiles(n uint) []WangTile {
    res := make([]WangTile, 0, n)
    row_size := uint(math.Sqrt(float64(n)))
    tiles := b.firstRow(row_size)
    res = append(res, tiles...)

    for i := uint(1); i < row_size; i++ {
        next := b.nextRow(tiles)
        res = append(res, next...)
        tiles = next
    }

    return res
}

func (b *Builder) Tiles(n uint) *Builder {
    b.tiles = b.genTiles(n)
    return b
}

func (b *Builder) Build() canvas.Canvas {
    fl := float64(len(b.tiles))
    n := math.Sqrt(fl)
    part := float32(1.0 / n)

    wg := sync.WaitGroup{}
    for i, tile := range b.tiles {
        i := i
        tile := tile
        
        wg.Add(1)

        go func() {
            xoff := part * float32(i % int(n))
            yoff := part * float32(math.Floor(float64(i) / n))
            coords := [4]Coord{
                {xoff, 1 - yoff},
                {xoff + part, 1 - yoff},
                {xoff + part, 1 - yoff - part},
                {xoff, 1 - yoff - part},
            }
            for _, t := range at(tile, coords) {
                b.canvas.DrawTriangle(t)
            }
            wg.Done()
        }()
    }

    wg.Wait()
    return b.canvas
}

func NewBuilder(c canvas.Canvas) Builder {
    return Builder{
        canvas: c,
    }
}
