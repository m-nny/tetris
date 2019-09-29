package env

import (
	"fmt"
	"math/rand"
)

type environment struct {
	board                   Board
	currentShape, nextShape Shape
	shapeX, shapeY          int
	score                   int
}

func NewEnvironment(seed int64) environment {
	rand.Seed(seed)
	env := environment{}
	env.currentShape, env.nextShape = getRandomShape(), getRandomShape()
	return env
}

func (env *environment) Render() {
	fmt.Println("#==========#")
	board, err := env.board.fit(env.currentShape, env.shapeX, env.shapeY)
	if err != nil {
		board = &env.board
	}
	for _, row := range board {
		fmt.Print("|")
		for _, cell := range row {
			if cell == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%v", cell)
			}
		}
		fmt.Println("|")
	}
	fmt.Println("#==========#")
	for _, row := range env.currentShape {
		fmt.Print("|")
		for _, cell := range row {
			fmt.Printf("%v", cell)
		}
		fmt.Println("|")
	}
	fmt.Println()
}

func (env *environment) Update() error {
	if err := env.board.canFit(env.currentShape, env.shapeX, env.shapeY); err != nil {
		return errGameOver()
	}
	if err := env.board.canFit(env.currentShape, env.shapeX, env.shapeY+1); err != nil {
		newBoard, _ := env.board.fit(env.currentShape, env.shapeX, env.shapeY)
		env.board = *newBoard

		env.currentShape, env.nextShape = env.nextShape, getRandomShape()
		env.shapeX, env.shapeY = 0, 0
	}
	env.shapeY++
	return nil
}

func (env *environment) MoveRight() error {
	if err := env.board.canFit(env.currentShape, env.shapeX+1, env.shapeY); err != nil {
		return errMove("right")
	}
	env.shapeX++
	return nil
}
