package roguelike

import "github.com/hajimehoshi/ebiten/v2"

type tile struct {
	X, Y    int
	Blocked bool
	Image   *ebiten.Image

	IsRevealed bool
}
