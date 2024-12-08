package SnakeGame

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	GameSpeed    = time.Second / 6
	ScreenWidth  = 640
	ScreenHeight = 480
	GridSize     = 20
)

var (
	Up              = Point{x: 0, y: -1}
	Down            = Point{x: 0, y: 1}
	Left            = Point{x: -1, y: 0}
	Right           = Point{x: 1, y: 0}
	MplusFaceSource *text.GoTextFaceSource
)

type Game struct {
	snake      []Point
	food       Point
	direction  Point
	lastUpdate time.Time
	gameOver   bool
}

func CreateNewGame() *Game {
	return &Game{
		snake: []Point{{
			x: ScreenWidth / GridSize / 2,
			y: ScreenHeight / GridSize / 2,
		}},
		direction: Point{
			x: 1,
			y: 0,
		},
		lastUpdate: time.Now(),
		food: Point{
			x: rand.Intn(ScreenWidth / GridSize),
			y: rand.Intn(ScreenHeight / GridSize),
		},
		gameOver: false,
	}
}

func (g *Game) spawnFood() {
	g.food = Point{
		x: rand.Intn(ScreenWidth / GridSize),
		y: rand.Intn(ScreenHeight / GridSize),
	}
}

func (g *Game) updateSnake(snake *[]Point, dir Point) {
	head := (*snake)[0]
	newHead := Point{
		x: head.x + dir.x,
		y: head.y + dir.y,
	}

	if g.isCollition(newHead, *snake) {
		g.gameOver = true
		GlobalAudioManager.StopMusic()
		GlobalAudioManager.PlaySound("game-over")
	}

	if g.food == newHead {
		*snake = append([]Point{newHead}, *snake...)
		GlobalAudioManager.PlaySound("food")
		g.spawnFood()
	} else {
		*snake = append([]Point{newHead}, (*snake)[:len(*snake)-1]...)
	}
}

func (g Game) isCollition(head Point, snake []Point) bool {
	if head.x < 0 || head.y < 0 {
		return true
	}
	if head.x >= ScreenWidth/GridSize || head.y >= ScreenHeight/GridSize {
		return true
	}

	for _, p := range snake {
		if p == head {
			return true
		}
	}

	return false
}

func (g *Game) handleKeyStroke() {
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		if g.gameOver {
			g.snake = []Point{{
				x: ScreenWidth / GridSize / 2,
				y: ScreenHeight / GridSize / 2,
			}}
			g.spawnFood()
			GlobalAudioManager.PlayMusic("music")
			g.gameOver = false
		}
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if len(g.snake) > 1 && g.direction == Down {
			return
		}
		g.direction = Up
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if len(g.snake) > 1 && g.direction == Up {
			return
		}
		g.direction = Down
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if len(g.snake) > 1 && g.direction == Right {
			return
		}
		g.direction = Left
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if len(g.snake) > 1 && g.direction == Left {
			return
		}
		g.direction = Right
		return
	}
}

func (g *Game) drawImage(screen *ebiten.Image, image *ebiten.Image, x, y int) error {
	if image == nil {
		return fmt.Errorf("image is nil at (%d, %d)", x, y)
	}

	options := &ebiten.DrawImageOptions{}

	imageWidth := float64(image.Bounds().Dx())
	imageHeight := float64(image.Bounds().Dy())
	scaleX := float64(GridSize) / imageWidth
	scaleY := float64(GridSize) / imageHeight
	options.GeoM.Scale(scaleX, scaleY)

	options.GeoM.Translate(
		float64(x*GridSize),
		float64(y*GridSize),
	)

	screen.DrawImage(image, options)
	return nil
}

func (g *Game) Update() error {
	g.handleKeyStroke()
	if g.gameOver {
		return nil
	}
	if time.Since(g.lastUpdate) < GameSpeed {
		return nil
	}
	g.lastUpdate = time.Now()
	g.updateSnake(&g.snake, g.direction)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for y := 0; y < ScreenHeight/GridSize; y++ {
		for x := 0; x < ScreenWidth/GridSize; x++ {
			g.drawImage(screen, Tile, x, y)
		}
	}
	for _, p := range g.snake {
		vector.DrawFilledRect(
			screen,
			float32(p.x*GridSize),
			float32(p.y*GridSize),
			GridSize,
			GridSize,
			color.White,
			true,
		)
	}

	g.drawImage(screen, Food, g.food.x, g.food.y)

	if g.gameOver {
		face := &text.GoTextFace{
			Source: MplusFaceSource,
			Size:   48,
		}
		w, h := text.Measure("Game Over!", face, face.Size)
		vector.DrawFilledRect(
			screen,
			float32((ScreenWidth-w)/2-GridSize),
			float32((ScreenHeight-h)/2-GridSize),
			float32(w+GridSize*2),
			float32(h+GridSize*2),
			color.Black,
			true,
		)
		options := &text.DrawOptions{}
		options.GeoM.Translate((ScreenWidth-w)/2, (ScreenHeight-h)/2)
		text.Draw(screen, "Game Over!", face, options)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
