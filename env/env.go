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
	board                   Board
	currentShape, nextShape Shape
	shapeX, shapeY          int
	score                   int
	timestamp               int
	rand                    *rand.Rand
}

// NewEnvironment initialize new game Environment
func NewEnvironment(seed int64) Environment {
	env := Environment{rand: getRand(seed)}
	env.currentShape, env.nextShape = getRandomShape(env.rand), getRandomShape(env.rand)
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
	newBoard, squashedLines := newBoard.squash()
	env.board = *newBoard

	env.score += scorePerLine[squashedLines]

	env.currentShape, env.nextShape = env.nextShape, getRandomShape(env.rand)
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

func getRand(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}
