package roguelike

import (
	"fmt"

	"github.com/bytearena/ecs"
)

func attackSystem(g *game, attackerPosition *position, defenderPosition *position) {
	var attacker *ecs.QueryResult
	var defender *ecs.QueryResult

	//Get the attacker and defender if either is a player
	for _, player := range g.World.Query(g.WorldTags[playersTag]) {
		pos := player.Components[positionComponent].(*position)
		if pos.IsEqual(attackerPosition) {
			//This is the attacker
			attacker = player
		} else if pos.IsEqual(defenderPosition) {
			//This is the defender
			defender = player
		}
	}

	// Get the attacker and defender if either is a monster
	for _, monster := range g.World.Query(g.WorldTags[monsterTag]) {
		pos := monster.Components[positionComponent].(*position)
		if pos.IsEqual(attackerPosition) {
			//This is the attacker
			attacker = monster
		} else if pos.IsEqual(defenderPosition) {
			//This is the defender
			defender = monster
		}
	}

	//If we somehow don't have an attacker or defender, just leave
	if attacker == nil || defender == nil {
		return
	}

	//Grab the required information
	attackerWeapon := attacker.Components[meleeWeaponComponent].(*meleeWeapon)
	attackerName := attacker.Components[nameComponent].(*name).Label
	attackerMessage := attacker.Components[messageComponent].(*message)

	defenderArmor := defender.Components[armorComponent].(*armor)
	defenderHealth := defender.Components[healthComponent].(*health)
	defenderMessage := defender.Components[messageComponent].(*message)
	defenderName := defender.Components[nameComponent].(*name).Label

	//Roll a d10 to hit
	toHitRoll := getDiceRoll(10)

	if toHitRoll+attackerWeapon.ToHitBonus > defenderArmor.ArmorClass {
		//It's a hit!
		damageRoll := getRandomIntBetween(attackerWeapon.MinimumDamage, attackerWeapon.MaximumDamage)

		damageDone := damageRoll - defenderArmor.Defense
		//Let's not have the weapon heal the defender
		if damageDone < 0 {
			damageDone = 0
		}
		defenderHealth.CurrentHealth -= damageDone

		attackerMessage.AttackMessage = fmt.Sprintf(
			"%s swings %s at %s and hits for %d health.\n",
			attackerName, attackerWeapon.Name, defenderName, damageDone,
		)

		if defenderHealth.CurrentHealth <= 0 {
			defenderMessage.DeadMessage = fmt.Sprintf("%s has died!\n", defenderName)
			if defenderName == "Player" {
				defenderMessage.GameStateMessage = "Game Over!\n"
				g.Turn = gameOver
			}
			g.World.DisposeEntity(defender.Entity)
		}
	} else {
		attackerMessage.AttackMessage = fmt.Sprintf("%s swings %s at %s and misses.\n", attackerName, attackerWeapon.Name, defenderName)
	}
}
