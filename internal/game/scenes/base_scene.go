package scenes

import (
	"image"

	"github.com/ebitenui/ebitenui"
	"github.com/joelschutz/stagehand"
	darkfalltypes "github.com/mavrw/ravenfall-proto/pkg/darkfall_types"
)

type BaseScene struct {
	// your scene fields
	sm     *stagehand.SceneManager[darkfalltypes.DarkfallGameState]
	ui     *ebitenui.UI
	state  darkfalltypes.DarkfallGameState
	bounds image.Rectangle
}

func (s *BaseScene) Load(state darkfalltypes.DarkfallGameState, manager stagehand.SceneController[darkfalltypes.DarkfallGameState]) {
	// your load code

	s.sm = manager.(*stagehand.SceneManager[darkfalltypes.DarkfallGameState]) // This type assertion is important
}

func (s *BaseScene) Unload() darkfalltypes.DarkfallGameState {
	// your unload code

	return s.state
}

func (s *BaseScene) Layout(w, h int) (int, int) {
	s.bounds = image.Rect(0, 0, w, h)
	return w, h
}

func (s *BaseScene) getUI() *ebitenui.UI {
	// TODO: Improve this. Should probably write very thin shim around ebitenui
	return s.ui
}
