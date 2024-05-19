package roguelike

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type game struct {
	GameMap   gameMap
	World     *ecs.Manager
	WorldTags map[string]ecs.Tag
}

func NewGame() *game {
	gameMap := newGameMap()
	world, tags := initializeWorld(gameMap.CurrentLevel)

	return &game{
		GameMap:   gameMap,
		World:     world,
		WorldTags: tags,
	}
}

func (g *game) Update() error {
	tryMovePlayer(g)
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	level := g.GameMap.CurrentLevel
	level.Draw(screen)

	processRenderables(g, level, screen)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
