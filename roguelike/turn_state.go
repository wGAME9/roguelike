package roguelike

type turnState int

const (
	beforePlayerAction turnState = iota
	playerTurn
	monsterTurn
	gameOver
)

func getNextState(state turnState) turnState {
	switch state {
	case beforePlayerAction:
		return playerTurn

	case playerTurn:
		return monsterTurn

	case monsterTurn:
		return beforePlayerAction

	case gameOver:
		return gameOver

	default:
		return playerTurn
	}
}
