package main

import (
	"log"

	SnakeGame "github.com/AbdelilahOu/Snake-game-go/snake-game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := SnakeGame.CreateNewGame()

	ebiten.SetWindowSize(SnakeGame.ScreenWidth, SnakeGame.ScreenHeight)
	ebiten.SetWindowTitle("Snake Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
