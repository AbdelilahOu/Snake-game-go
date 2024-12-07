package Resources

import (
	_ "embed"
)

//go:embed fonts/04B_30__.TTF
var PixelFontFace []byte

//go:embed images/apple.png
var Apple []byte

var (
	//go:embed sounds/gameover.mp3
	GameOverSound []byte
	//go:embed sounds/food.mp3
	FoodSound []byte
	//go:embed sounds/music.mp3
	MusicSound []byte
	//go:embed sounds/move.mp3
	MoveSound []byte
)
