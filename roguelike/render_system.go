package roguelike

import "github.com/hajimehoshi/ebiten/v2"

func processRenderables(game *game, level level, screen *ebiten.Image) {
	for _, result := range game.World.Query(game.WorldTags[renderablesTag]) {
		pos := result.Components[positionComponent].(*position)
		img := result.Components[renderableComponent].(*renderable).Image

		isVisible := level.PlayerVisible.IsVisible(pos.X, pos.Y)
		if isVisible {
			index := level.getIndexFromCoords(pos.X, pos.Y)
			tile := level.Tiles[index]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.X), float64(tile.Y))
			screen.DrawImage(img, op)
		}

	}
}
