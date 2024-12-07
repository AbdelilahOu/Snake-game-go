package main

import (
	"bytes"
	"log"

	SnakeGame "github.com/AbdelilahOu/Snake-game-go/snake-game"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func main() {
	// load font
	s, err := text.NewGoTextFaceSource(
		bytes.NewReader(
			fonts.MPlus1pRegular_ttf,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	SnakeGame.MplusFaceSource = s
	game := SnakeGame.CreateNewGame()

	ebiten.SetWindowSize(SnakeGame.ScreenWidth, SnakeGame.ScreenHeight)
	ebiten.SetWindowTitle("Snake Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
