package env

import "fmt"

var ok = error(nil)

func errBounds() error {
	return fmt.Errorf("env: Out of bounds")
}

func errFilled() error {
	return fmt.Errorf("env: Cell is already filled with another shape")
}

func errGameOver() error {
	return fmt.Errorf("env: Game is over")
}

func errMove(dir string) error {
	return fmt.Errorf("env: Cannot move to %v", dir)
}
