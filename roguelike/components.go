package roguelike

import "github.com/hajimehoshi/ebiten/v2"

type player struct{}

type monster struct{}

type position struct {
	X int
	Y int
}

type renderable struct {
	Image *ebiten.Image
}

type moveable struct{}
