package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/wGAME9/roguelike/roguelike"
)

func main() {
	game := roguelike.NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Tower")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
