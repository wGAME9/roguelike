package roguelike

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type player struct{}

type monster struct {
	Name string
}

type position struct {
	X int
	Y int
}

func (p *position) GetManhattanDistance(other *position) int {
	xDist := math.Abs(float64(p.X - other.X))
	yDist := math.Abs(float64(p.Y - other.Y))
	return int(xDist) + int(yDist)
}

type renderable struct {
	Image *ebiten.Image
}

type moveable struct{}
