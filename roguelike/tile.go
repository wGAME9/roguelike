package roguelike

import "github.com/hajimehoshi/ebiten/v2"

type typeOfTile int

const (
	WALL typeOfTile = iota
	FLOOR
)

type tile struct {
	TypeOfTile typeOfTile
	X, Y       int
	Blocked    bool
	Image      *ebiten.Image

	IsRevealed bool
}
