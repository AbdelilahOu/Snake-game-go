package SnakeGame

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	GameSpeed    = time.Second / 6
	ScreenWidth  = 640
	ScreenHeight = 480
	GridSize     = 20
)

var (
	Up    = Point{x: 0, y: -1}
	Down  = Point{x: 0, y: 1}
	Left  = Point{x: -1, y: 0}
	Right = Point{x: 1, y: 0}
)

type Game struct {
	snake      []Point
	food       Point
	direction  Point
	lastUpdate time.Time
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

	if g.food == newHead {
		*snake = append([]Point{newHead}, *snake...)
		g.spawnFood()
	} else {
		*snake = append([]Point{newHead}, (*snake)[:len(*snake)-1]...)
	}
}

func (g *Game) updateDirection() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.direction = Up
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.direction = Down
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.direction = Left
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.direction = Right
		return
	}
}

func (g *Game) Update() error {
	g.updateDirection()
	if time.Since(g.lastUpdate) < GameSpeed {
		return nil
	}
	g.lastUpdate = time.Now()
	g.updateSnake(&g.snake, g.direction)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
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

		vector.DrawFilledRect(
			screen,
			float32(g.food.x*GridSize),
			float32(g.food.y*GridSize),
			GridSize,
			GridSize,
			color.RGBA{255, 0, 0, 255},
			true,
		)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
