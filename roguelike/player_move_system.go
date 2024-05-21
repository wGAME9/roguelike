package roguelike

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func tryMovePlayer(game *game) {
	turnTaken := false

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		turnTaken = true
	}

	dx, dy := 0, 0

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		dy = -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		dy = 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		dx = -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		dx = 1
	}

	level := game.GameMap.CurrentLevel

	player := game.WorldTags[playersTag]

	for _, result := range game.World.Query(player) {
		pos := result.Components[positionComponent].(*position)
		index := level.getIndexFromCoords(pos.X+dx, pos.Y+dy)

		tile := level.Tiles[index]
		if tile.Blocked {
			if dx != 0 || dy != 0 {
				if level.Tiles[index].TypeOfTile != WALL {
					//Its a tile with a monster -- Fight it
					monsterPosition := position{X: pos.X + dx, Y: pos.Y + dy}
					attackSystem(game, pos, &monsterPosition)
				}
			}
			break
		}

		level.Tiles[level.getIndexFromCoords(pos.X, pos.Y)].Blocked = false

		pos.X += dx
		pos.Y += dy
		// making monster is blocked from moving to where player is standing
		level.Tiles[index].Blocked = true
		level.PlayerVisible.Compute(level, pos.X, pos.Y, 8)
	}

	if dx != 0 || dy != 0 || turnTaken {
		game.Turn = getNextState(game.Turn)
		game.TurnCounter = 0
	}
}
