package roguelike

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	wallImage   *ebiten.Image
	floorImage  *ebiten.Image
	playerImage *ebiten.Image
	skellyImage *ebiten.Image
	orcImage    *ebiten.Image
	logImg      *ebiten.Image
)

func init() {
	var err error
	wallImage, _, err = ebitenutil.NewImageFromFile("assets/wall.png")
	if err != nil {
		log.Fatal(err)
	}

	floorImage, _, err = ebitenutil.NewImageFromFile("assets/floor.png")
	if err != nil {
		log.Fatal(err)
	}

	playerImage, _, err = ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}

	skellyImage, _, err = ebitenutil.NewImageFromFile("assets/skelly.png")
	if err != nil {
		log.Fatal(err)
	}

	orcImage, _, err = ebitenutil.NewImageFromFile("assets/orc.png")
	if err != nil {
		log.Fatal(err)
	}

	logImg, _, err = ebitenutil.NewImageFromFile("assets/UIPanel.png")
	if err != nil {
		log.Fatal(err)
	}
}
