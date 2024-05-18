package roguelike

type turnState int

const (
	beforePlayerAction turnState = iota
	playerTurn
	monsterTurn
)

func getNextState(state turnState) turnState {
	switch state {
	case beforePlayerAction:
		return playerTurn

	case playerTurn:
		return monsterTurn

	case monsterTurn:
		return beforePlayerAction

	default:
		return playerTurn
	}
}
