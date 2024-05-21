package roguelike

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	lastText []string = make([]string, 0, 5)
)

func processUserLog(g *game, screen *ebiten.Image) {
	uiLocation := (numTilesY - uiHeight) * tileHeight
	var fontX = 16
	var fontY = uiLocation + 24
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(0.), float64(uiLocation))
	screen.DrawImage(logImg, op)

	tmpMessages := make([]string, 0, 5)
	anyMessages := false

	for _, m := range g.World.Query(g.WorldTags[messageTag]) {
		messages := m.Components[messageComponent].(*message)
		if messages.AttackMessage != "" {
			tmpMessages = append(tmpMessages, messages.AttackMessage)
			anyMessages = true
			messages.AttackMessage = ""
		}
	}
	for _, m := range g.World.Query(g.WorldTags[messageTag]) {
		messages := m.Components[messageComponent].(*message)
		if messages.DeadMessage != "" {
			tmpMessages = append(tmpMessages, messages.DeadMessage)
			anyMessages = true
			messages.DeadMessage = ""
			g.World.DisposeEntity(m.Entity)
		}
		if messages.GameStateMessage != "" {
			tmpMessages = append(tmpMessages, messages.GameStateMessage)
			anyMessages = true
			//No need to clear, it's all over
		}

	}
	if anyMessages {
		lastText = tmpMessages
	}
	for _, msg := range lastText {
		if msg != "" {
			op := &text.DrawOptions{}
			op.GeoM.Translate(float64(fontX), float64(fontY))
			op.ColorScale.ScaleWithColor(color.White)

			text.Draw(screen, msg, &text.GoTextFace{
				Source: mplusFaceSource,
				Size:   normalFontSize,
			}, op)
			fontY += 16
		}
	}
}
