package roguelike

import (
	"github.com/bytearena/ecs"
)

const (
	playersTag     = "players"
	renderablesTag = "renderables"
)

var (
	positionComponent   *ecs.Component
	renderableComponent *ecs.Component
)

func initializeWorld(startingLevel level) (*ecs.Manager, map[string]ecs.Tag) {
	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.center()

	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	playerComponent := manager.NewComponent()
	moveableComponent := manager.NewComponent()

	positionComponent = manager.NewComponent()
	renderableComponent = manager.NewComponent()

	manager.NewEntity().
		AddComponent(playerComponent, player{}).
		AddComponent(moveableComponent, moveable{}).
		AddComponent(renderableComponent, &renderable{
			Image: playerImage,
		}).
		AddComponent(positionComponent, &position{X: x, Y: y})

	players := ecs.BuildTag(playerComponent, positionComponent)
	tags[playersTag] = players

	renderables := ecs.BuildTag(renderableComponent, positionComponent)
	tags[renderablesTag] = renderables

	return manager, tags
}
