package main

import (
	"bytes"
	"log"

	Resources "github.com/AbdelilahOu/Snake-game-go/resources"
	SnakeGame "github.com/AbdelilahOu/Snake-game-go/snake-game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func main() {
	// load font
	s, err := text.NewGoTextFaceSource(
		bytes.NewReader(
			Resources.PixelFontFace,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	SnakeGame.MplusFaceSource = s
	game := SnakeGame.CreateNewGame()

	ebiten.SetWindowSize(SnakeGame.ScreenWidth, SnakeGame.ScreenHeight)
	ebiten.SetWindowTitle("Snake Game")
	SnakeGame.GlobalAudioManager.PlayMusic("music")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
