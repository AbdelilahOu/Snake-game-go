# Snake Game

![Preview of the Snake Game](/resources/preview.png)

## Description

This is a classic snake game built using the Go programming language and the Ebiten game engine. The player controls a snake that grows in length by eating food. The game ends when the snake collides with the walls or its own body.

## Getting Started

1. Make sure you have Go and the Ebiten library installed on your system.
2. Clone the repository or download the project files.
   ```
   git clone https://github.com/AbdelilahOu/Snake-game-go.git
   ```
3. Navigate to the project directory and run the game using the following command:
   ```
   go run main.go
   ```

## File Structure

```
go-snake-game/
├── resources/
│   ├── preview.png
│   ├── resources.go
│   ├── fonts/
│   │   └── As_028.3.0.TTF
│   ├── images/
│   │   ├── apple.png
│   │   └── tile.png
│   └── sounds/
│       ├── food.mp3
│       ├── move.mp3
│       ├── gameover.mp3
│       └── music.mp3
├── snake-game/
│   ├── background.go
│   ├── food.go
│   ├── game.go
│   ├── snake.go
│   └── sound.go
├── go.mod
├── go.sum
├── main.go
└── README.md
```
