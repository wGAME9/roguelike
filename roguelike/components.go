package roguelike

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type name struct {
	Label string
}

type player struct{}

type monster struct {
}

type health struct {
	MaxHealth     int
	CurrentHealth int
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

func (p *position) IsEqual(other *position) bool {
	return (p.X == other.X && p.Y == other.Y)
}

type renderable struct {
	Image *ebiten.Image
}

type moveable struct{}

type meleeWeapon struct {
	Name          string
	MinimumDamage int
	MaximumDamage int
	ToHitBonus    int
}

type armor struct {
	Name       string
	Defense    int
	ArmorClass int
}
