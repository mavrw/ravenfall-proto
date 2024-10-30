package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mavrw/ravenfall-proto/internal/game"
)

func main() {
	game := game.NewDarkfallGame()

	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle(game.Name)

	if err := game.Start(); err != nil {
		fmt.Println(err)
	}
}
