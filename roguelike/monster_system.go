package roguelike

import (
	"github.com/norendren/go-fov/fov"
)

func updateMonster(game *game) {
	l := game.GameMap.CurrentLevel
	playerPosition := position{}

	for _, plr := range game.World.Query(game.WorldTags[playersTag]) {
		pos := plr.Components[positionComponent].(*position)
		playerPosition.X = pos.X
		playerPosition.Y = pos.Y
	}

	for _, result := range game.World.Query(game.WorldTags[monsterTag]) {
		pos := result.Components[positionComponent].(*position)
		monsterSees := fov.New()
		monsterSees.Compute(l, pos.X, pos.Y, 8)

		monsterCanSeesPlayer := monsterSees.IsVisible(playerPosition.X, playerPosition.Y)
		if monsterCanSeesPlayer {
			if pos.GetManhattanDistance(&playerPosition) == 1 {
				attackSystem(game, pos, &playerPosition)
				continue
			}

			astar := AStar{}
			path := astar.GetPath(l, pos, &playerPosition)
			if len(path) > 1 {
				nextTile := l.Tiles[l.getIndexFromCoords(path[1].X, path[1].Y)]
				if !nextTile.Blocked {
					l.Tiles[l.getIndexFromCoords(pos.X, pos.Y)].Blocked = false
					pos.X = path[1].X
					pos.Y = path[1].Y
					nextTile.Blocked = true
				}
			}
		}
	}
	game.Turn = playerTurn
}
