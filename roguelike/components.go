package roguelike

import "github.com/hajimehoshi/ebiten/v2"

type player struct{}

type position struct {
	X int
	Y int
}

type renderable struct {
	Image *ebiten.Image
}

type moveable struct{}
