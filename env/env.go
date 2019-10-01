package env

import (
	"fmt"
	"math/rand"
)

var scorePerLine = map[int]int{
	0: 0,
	1: 40,
	2: 100,
	3: 300,
	4: 1200,
}

// Environment current Game representation
type Environment struct {
	Board                   Board
	currentShape, nextShape Shape
	ShapeX, ShapeY          int
	score                   int
	timestamp               int
	rand                    *rand.Rand
}

// NewEnvironment initialize new game Environment
func NewEnvironment(seed int64) Environment {
	env := Environment{rand: GetRand(seed)}
	env.currentShape, env.nextShape = getRandomShape(env.rand), getRandomShape(env.rand)
	return env
}

// Render game in terminal (stdout)
func (env *Environment) Render() {
	fmt.Printf("#%v\n", env.timestamp)
	fmt.Println("#==========#")
	board, err := env.Board.fit(env.currentShape, env.ShapeX, env.ShapeY)
	if err != nil {
		board = &env.Board
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
	if err := env.Board.canFit(env.currentShape, env.ShapeX, env.ShapeY); err != nil {
		return errGameOver()
	}
	if err := env.Board.canFit(env.currentShape, env.ShapeX, env.ShapeY+1); err == nil {
		env.ShapeY++
		return nil
	}
	newBoard, _ := env.Board.fit(env.currentShape, env.ShapeX, env.ShapeY)
	newBoard, squashedLines := newBoard.squash()
	env.Board = *newBoard

	env.score += scorePerLine[squashedLines]

	env.currentShape, env.nextShape = env.nextShape, getRandomShape(env.rand)
	env.ShapeX, env.ShapeY = 0, 0
	return nil
}

// MakeAction on environment
func (env *Environment) MakeAction(a Action) error {
	switch a {
	case MoveLeft:
		if err := env.Board.canFit(env.currentShape, env.ShapeX-1, env.ShapeY); err != nil {
			return errMove("left")
		}
		env.ShapeX--
	case MoveRight:
		if err := env.Board.canFit(env.currentShape, env.ShapeX+1, env.ShapeY); err != nil {
			return errMove("right")
		}
		env.ShapeX++
	}
	return nil
}
