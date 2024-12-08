package SnakeGame

import (
	"bytes"
	"log"

	Resources "github.com/AbdelilahOu/Snake-game-go/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Background struct {
	rendered bool
}

var Tile *ebiten.Image

func init() {
	var err error
	Tile, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(Resources.Tile))
	if err != nil {
		log.Fatal(err)
	}
}

func CreateNewBackground() *Background {
	return &Background{
		rendered: false,
	}
}