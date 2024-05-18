package roguelike

type gameMap struct {
	Dungeons     []dungeon
	CurrentLevel level
}

func newGameMap() gameMap {
	defaultDungeon := createDefaultDungeon()
	defaultLevel := defaultDungeon.Levels[0]
	return gameMap{
		Dungeons:     []dungeon{defaultDungeon},
		CurrentLevel: defaultLevel,
	}
}

func createDefaulLevel() level {
	return newLevel()
}

func createDefaultDungeon() dungeon {
	return dungeon{
		Name:   "default",
		Levels: []level{createDefaulLevel()},
	}
}
