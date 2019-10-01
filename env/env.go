package env

import (
	"fmt"
	"math/rand"
)

// Environment current Game representation
type Environment struct {
	board                   Board
	currentShape, nextShape Shape
	shapeX, shapeY          int
	score                   int
	timestamp               int
}

// NewEnvironment initialize new game Environment
func NewEnvironment(seed int64) Environment {
	rand.Seed(seed)
	env := Environment{}
	env.currentShape, env.nextShape = getRandomShape(), getRandomShape()
	return env
}

// Render game in terminal (stdout)
func (env *Environment) Render() {
	fmt.Printf("#%v\n", env.timestamp)
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
	fmt.Print("\n\n")
}

// update environment state
func (env *Environment) update() error {
	env.timestamp++
	if err := env.board.canFit(env.currentShape, env.shapeX, env.shapeY); err != nil {
		return errGameOver()
	}
	if err := env.board.canFit(env.currentShape, env.shapeX, env.shapeY+1); err == nil {
		env.shapeY++
		return nil
	}
	newBoard, _ := env.board.fit(env.currentShape, env.shapeX, env.shapeY)
	env.board = *newBoard

	env.score += 100

	env.currentShape, env.nextShape = env.nextShape, getRandomShape()
	env.shapeX, env.shapeY = 0, 0
	return nil
}

// MakeAction on environment
func (env *Environment) MakeAction(a Action) error {
	switch a {
	case MoveLeft:
		if err := env.board.canFit(env.currentShape, env.shapeX-1, env.shapeY); err != nil {
			return errMove("left")
		}
		env.shapeX--
	case MoveRight:
		if err := env.board.canFit(env.currentShape, env.shapeX+1, env.shapeY); err != nil {
			return errMove("right")
		}
		env.shapeX++
	}
	return nil
}
