package SnakeGame

import (
	"bytes"
	_ "image/png"
	"log"

	Resources "github.com/AbdelilahOu/Snake-game-go/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var Food *ebiten.Image

func init() {
	var err error
	Food, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(Resources.Apple))
	if err != nil {
		log.Fatal(err)
	}
}
