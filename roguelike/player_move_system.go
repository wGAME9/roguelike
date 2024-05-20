package roguelike

import "github.com/hajimehoshi/ebiten/v2"

func tryMovePlayer(game *game) {
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
			return
		}

		pos.X += dx
		pos.Y += dy
		level.PlayerVisible.Compute(level, pos.X, pos.Y, 8)
	}

	if dx != 0 || dy != 0 {
		game.Turn = getNextState(game.Turn)
		game.TurnCounter = 0
	}
}
