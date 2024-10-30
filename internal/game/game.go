package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joelschutz/stagehand"
	"github.com/mavrw/ravenfall-proto/internal/game/scenes"
	darkfalltypes "github.com/mavrw/ravenfall-proto/pkg/darkfall_types"
)

type DarkfallGame struct {
	// The name of the game
	Name         string
	ScreenWidth  int
	ScreenHeight int

	sceneManager *stagehand.SceneManager[darkfalltypes.DarkfallGameState]
}

func NewDarkfallGame() DarkfallGame {
	return DarkfallGame{
		Name:         "Darkfall",
		ScreenWidth:  1280,
		ScreenHeight: 720,
	}
}

func (g DarkfallGame) Start() error {
	mainMenu := &scenes.MainMenuScene{}
	state := darkfalltypes.DarkfallGameState{}
	g.sceneManager = stagehand.NewSceneManager(mainMenu, state)

	if err := ebiten.RunGame(g.sceneManager); err != nil {
		return err
	}

	return nil
}
