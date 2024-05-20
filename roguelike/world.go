package roguelike

import (
	"github.com/bytearena/ecs"
)

const (
	playersTag     = "players"
	renderablesTag = "renderables"
	monsterTag     = "monsters"
)

var (
	monsterComponent    *ecs.Component
	positionComponent   *ecs.Component
	renderableComponent *ecs.Component
)

func initializeWorld(startingLevel level) (*ecs.Manager, map[string]ecs.Tag) {
	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()

	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	playerComponent := manager.NewComponent()
	monsterComponent = manager.NewComponent()

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

	for _, room := range startingLevel.Rooms {
		if room.X1 != startingRoom.X1 {
			mX, mY := room.Center()

			manager.NewEntity().
				AddComponent(monsterComponent, &monster{
					Name: "Skelton",
				}).
				AddComponent(renderableComponent, &renderable{
					Image: skellyImage,
				}).
				AddComponent(positionComponent, &position{
					X: mX,
					Y: mY,
				})
		}
	}

	monsters := ecs.BuildTag(monsterComponent, positionComponent)
	tags["monsters"] = monsters

	return manager, tags
}
