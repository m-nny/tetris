package main

import (
	"fmt"

	"github.com/m-nny/tetris/env"
)

func main() {
	env := env.NewEnvironment(0)
	env.Render()
	for i := 0; i < 100; i++ {
		if i%5 == 0 {
			env.MoveRight()
		}
		err := env.Update()
		env.Render()
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	env.Render()
}
