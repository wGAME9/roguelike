package roguelike

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/norendren/go-fov/fov"
)

type level struct {
	Tiles         []tile
	Rooms         []rect
	PlayerVisible *fov.View
}

func newLevel() level {
	l := level{
		Rooms:         make([]rect, 0),
		PlayerVisible: fov.New(),
	}
	l.GenerateLevelTiles()

	return l
}

func (l *level) Draw(screen *ebiten.Image) {
	for x := range numTilesX {
		for y := range numTilesY {
			isVisible := l.PlayerVisible.IsVisible(x, y)
			tileIdx := l.getIndexFromCoords(x, y)
			tile := l.Tiles[tileIdx]
			if !isVisible && !tile.IsRevealed {
				continue
			}

			var clrM colorm.ColorM
			op := &colorm.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.X), float64(tile.Y))
			if !isVisible && tile.IsRevealed {
				// it's not current visible, but it's been seen before
				clrM.Scale(0.5, 0.5, 0.5, 1)
			}
			colorm.DrawImage(screen, tile.Image, clrM, op)

			l.Tiles[tileIdx].IsRevealed = true
		}
	}
}

func (l *level) GenerateLevelTiles() {
	MIN_SIZE := 6
	MAX_SIZE := 10
	MAX_ROOMS := 30

	if len(l.Tiles) == 0 {
		l.createTiles()
	}

	for i := 0; i < MAX_ROOMS; i++ {
		w := getRandomIntBetween(MIN_SIZE, MAX_SIZE)
		h := getRandomIntBetween(MIN_SIZE, MAX_SIZE)
		x := getDiceRoll(numTilesX - w - 1)
		y := getDiceRoll(numTilesY - h - 1)

		newRoom := newRectangle(x, y, w, h)
		canAddThisRoom := true
		for _, otherRoom := range l.Rooms {
			if newRoom.Intersect(otherRoom) {
				canAddThisRoom = false
				break
			}
		}

		if canAddThisRoom {
			l.createRooms(newRoom)
			hasPreviousRoom := len(l.Rooms) > 0
			if hasPreviousRoom {
				newX, newY := newRoom.Center()
				lastRoom := l.Rooms[len(l.Rooms)-1]
				prevX, prevY := lastRoom.Center()

				coinFlip := getDiceRoll(2)
				if coinFlip == 2 {
					l.createHorizontalTunnel(prevX, newX, prevY)
					l.createVerticalTunnel(prevY, newY, newX)
				} else {
					l.createHorizontalTunnel(prevX, newX, newY)
					l.createVerticalTunnel(prevY, newY, prevX)
				}
			}

			l.Rooms = append(l.Rooms, newRoom)
		}
	}
}

func (l *level) createTiles() {
	tiles := make([]tile, numTilesX*numTilesY)

	for xIdx := range numTilesX {
		for yIdx := range numTilesY {
			placeIndx := l.getIndexFromCoords(xIdx, yIdx)
			tiles[placeIndx] = tile{
				X:          xIdx * tileWidth,
				Y:          yIdx * tileHeight,
				Blocked:    true,
				Image:      wallImage,
				IsRevealed: false,
				TypeOfTile: WALL,
			}
		}
	}

	l.Tiles = tiles
}

func (l *level) createRooms(room rect) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := l.getIndexFromCoords(x, y)
			l.Tiles[index].Blocked = false
			l.Tiles[index].TypeOfTile = FLOOR
			l.Tiles[index].Image = floorImage
		}
	}
}

func (l *level) createHorizontalTunnel(x1, x2, y int) {
	minX := min(x1, x2)
	maxX := max(x1, x2)
	for x := minX; x < maxX+1; x++ {
		index := l.getIndexFromCoords(x, y)
		if index > 0 && index < numTilesX*numTilesY {
			l.Tiles[index].Blocked = false
			l.Tiles[index].TypeOfTile = FLOOR
			l.Tiles[index].Image = floorImage
		}
	}
}

func (l *level) createVerticalTunnel(y1, y2, x int) {
	minY := min(y1, y2)
	maxY := max(y1, y2)
	for y := minY; y < maxY+1; y++ {
		index := l.getIndexFromCoords(x, y)
		if index > 0 && index < numTilesX*numTilesY {
			l.Tiles[index].Blocked = false
			l.Tiles[index].TypeOfTile = FLOOR
			l.Tiles[index].Image = floorImage
		}
	}
}

func (l level) InBounds(x, y int) bool {
	xIsInBound := 0 < x && x < numTilesX
	yIsInBound := 0 < y && y < numTilesY

	return xIsInBound && yIsInBound
}

func (l level) IsOpaque(x, y int) bool {
	idx := l.getIndexFromCoords(x, y)
	return l.Tiles[idx].TypeOfTile == WALL
}

func (l *level) getIndexFromCoords(x, y int) int {
	return y*numTilesX + x
}
