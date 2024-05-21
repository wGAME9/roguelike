package roguelike

import (
	"github.com/bytearena/ecs"
)

const (
	playersTag     = "players"
	renderablesTag = "renderables"
	monsterTag     = "monsters"
	messageTag     = "messages"
)

var (
	monsterComponent    *ecs.Component
	positionComponent   *ecs.Component
	renderableComponent *ecs.Component

	healthComponent      *ecs.Component
	meleeWeaponComponent *ecs.Component
	armorComponent       *ecs.Component
	nameComponent        *ecs.Component

	messageComponent *ecs.Component
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

	healthComponent = manager.NewComponent()
	meleeWeaponComponent = manager.NewComponent()
	armorComponent = manager.NewComponent()
	nameComponent = manager.NewComponent()

	messageComponent = manager.NewComponent()

	manager.NewEntity().
		AddComponent(playerComponent, player{}).
		AddComponent(moveableComponent, moveable{}).
		AddComponent(renderableComponent, &renderable{
			Image: playerImage,
		}).
		AddComponent(positionComponent, &position{X: x, Y: y}).
		AddComponent(healthComponent, &health{
			MaxHealth:     30,
			CurrentHealth: 30,
		}).
		AddComponent(meleeWeaponComponent, &meleeWeapon{
			Name:          "Battle Axe",
			MinimumDamage: 10,
			MaximumDamage: 20,
			ToHitBonus:    3,
		}).
		AddComponent(armorComponent, &armor{
			Name:       "Plate Armor",
			Defense:    15,
			ArmorClass: 18,
		}).
		AddComponent(nameComponent, &name{
			Label: "Player",
		}).
		AddComponent(messageComponent, &message{
			AttackMessage:    "",
			DeadMessage:      "",
			GameStateMessage: "",
		})

	renderables := ecs.BuildTag(renderableComponent, positionComponent)
	tags[renderablesTag] = renderables

	for _, room := range startingLevel.Rooms {
		if room.X1 != startingRoom.X1 {
			mX, mY := room.Center()

			manager.NewEntity().
				AddComponent(monsterComponent, &monster{}).
				AddComponent(renderableComponent, &renderable{
					Image: skellyImage,
				}).
				AddComponent(positionComponent, &position{
					X: mX,
					Y: mY,
				}).
				AddComponent(healthComponent, &health{
					MaxHealth:     10,
					CurrentHealth: 10,
				}).
				AddComponent(meleeWeaponComponent, &meleeWeapon{
					Name:          "Short Sword",
					MinimumDamage: 2,
					MaximumDamage: 6,
					ToHitBonus:    0,
				}).
				AddComponent(armorComponent, &armor{
					Name:       "Bone",
					Defense:    3,
					ArmorClass: 4,
				}).
				AddComponent(nameComponent, &name{
					Label: "Skeleton",
				}).
				AddComponent(messageComponent, &message{
					AttackMessage:    "",
					DeadMessage:      "",
					GameStateMessage: "",
				})
		}
	}

	players := ecs.BuildTag(
		playerComponent,
		positionComponent,
		healthComponent,
		meleeWeaponComponent,
		armorComponent,
		nameComponent,
		messageComponent,
	)
	tags[playersTag] = players

	monsters := ecs.BuildTag(
		monsterComponent,
		positionComponent,
		healthComponent,
		meleeWeaponComponent,
		armorComponent,
		nameComponent,
		messageComponent,
	)
	tags[monsterTag] = monsters

	messengers := ecs.BuildTag(messageComponent)
	tags[messageTag] = messengers

	return manager, tags
}
