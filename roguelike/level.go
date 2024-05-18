package roguelike

import "github.com/hajimehoshi/ebiten/v2"

type level struct {
	Tiles []tile
}

func newLevel() level {
	l := level{}
	l.createTiles()

	return l
}

func (l *level) Draw(screen *ebiten.Image) {
	for x := range numTilesX {
		for y := range numTilesY {
			tile := l.Tiles[l.getIndexFromCoords(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.X), float64(tile.Y))
			screen.DrawImage(tile.Image, op)
		}
	}
}

func (l *level) createTiles() {
	tiles := make([]tile, numTilesX*numTilesY)

	for xIdx := range numTilesX {
		for yIdx := range numTilesY {
			placeIndx := l.getIndexFromCoords(xIdx, yIdx)

			xIsOnEdge := xIdx == 0 || xIdx == numTilesX-1
			yIsOnEdge := yIdx == 0 || yIdx == numTilesY-1

			if xIsOnEdge || yIsOnEdge {
				tiles[placeIndx] = tile{
					X:       xIdx * tileWidth,
					Y:       yIdx * tileHeight,
					Blocked: true,
					Image:   wallImage,
				}

			} else {
				tiles[placeIndx] = tile{
					X:       xIdx * tileWidth,
					Y:       yIdx * tileHeight,
					Blocked: false,
					Image:   floorImage,
				}
			}
		}
	}

	l.Tiles = tiles
}

func (l *level) getIndexFromCoords(x, y int) int {
	return y*numTilesX + x
}
