package roguelike

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func processHUD(g *game, screen *ebiten.Image) {
	uiY := (numTilesY - uiHeight) * tileHeight
	uiX := (numTilesX * tileWidth) / 2
	var fontX = uiX + 16
	var fontY = uiY + 24

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(uiX), float64(uiY))
	screen.DrawImage(logImage, op)

	for _, player := range g.World.Query(g.WorldTags[playersTag]) {
		h := player.Components[healthComponent].(*health)
		healthText := fmt.Sprintf("Health: %d / %d", h.CurrentHealth, h.MaxHealth)

		drawTextOp := &text.DrawOptions{}
		drawTextOp.ColorScale.ScaleWithColor(color.White)
		drawTextOp.GeoM.Translate(float64(fontX), float64(fontY))
		text.Draw(
			screen,
			healthText,
			&text.GoTextFace{
				Source: mplusFaceSource,
				Size:   normalFontSize,
			},
			drawTextOp,
		)

		fontY += 16
		ac := player.Components[armorComponent].(*armor)
		acText := fmt.Sprintf("Armor Class: %d", ac.ArmorClass)

		drawTextOp = &text.DrawOptions{}
		drawTextOp.ColorScale.ScaleWithColor(color.White)
		drawTextOp.GeoM.Translate(float64(fontX), float64(fontY))
		text.Draw(
			screen,
			acText,
			&text.GoTextFace{
				Source: mplusFaceSource,
				Size:   normalFontSize,
			},
			drawTextOp,
		)

		fontY += 16
		defText := fmt.Sprintf("Defense: %d", ac.Defense)

		drawTextOp = &text.DrawOptions{}
		drawTextOp.ColorScale.ScaleWithColor(color.White)
		drawTextOp.GeoM.Translate(float64(fontX), float64(fontY))
		text.Draw(
			screen,
			defText,
			&text.GoTextFace{
				Source: mplusFaceSource,
				Size:   normalFontSize,
			},
			drawTextOp,
		)

		fontY += 16
		wpn := player.Components[meleeWeaponComponent].(*meleeWeapon)
		dmg := fmt.Sprintf("Damage: %d - %d", wpn.MinimumDamage, wpn.MaximumDamage)

		drawTextOp = &text.DrawOptions{}
		drawTextOp.ColorScale.ScaleWithColor(color.White)
		drawTextOp.GeoM.Translate(float64(fontX), float64(fontY))
		text.Draw(
			screen,
			dmg,
			&text.GoTextFace{
				Source: mplusFaceSource,
				Size:   normalFontSize,
			},
			drawTextOp,
		)

		fontY += 16
		bonus := fmt.Sprintf("To Hit Bonus: %d", wpn.ToHitBonus)

		drawTextOp = &text.DrawOptions{}
		drawTextOp.ColorScale.ScaleWithColor(color.White)
		drawTextOp.GeoM.Translate(float64(fontX), float64(fontY))
		text.Draw(
			screen,
			bonus,
			&text.GoTextFace{
				Source: mplusFaceSource,
				Size:   normalFontSize,
			},
			drawTextOp,
		)
	}
}
