package env

import "testing"

func TestNewEnvironment(t *testing.T) {
	seed := int64(42)
	env := NewEnvironment(seed)
	sameSeedEnv := NewEnvironment(seed)

	if !env.equal(&sameSeedEnv) {
		t.Errorf("NewEnvironment returned different results, with same seed")
	}
}

func TestEnvUpdate(t *testing.T) {
	env := NewEnvironment(42)
	if got := env.update(); got != ok {
		t.Errorf("env.Update() threw error %v", got)
	}
	if env.ShapeY != 1 {
		t.Errorf("env.Update() did not moved on to next line")
	}

	// On last line
	env.currentShape, env.ShapeY = iShape, BoardHeight-1
	if got := env.update(); got != ok {
		t.Errorf("env.Update() threw error %v", got)
	}
	if env.ShapeY != 0 {
		t.Errorf("env.Update() did not moved on to next shape")
	}

	env.Board[0][0] = 1
	if got := env.update(); got == ok {
		t.Errorf("env.Update() did not threw Game Over error")
	}
}

func TestRender(t *testing.T) {
	// TODO: THIS SHOULD BE REPLACED WITH PROPER TEST
	env := NewEnvironment(42)
	env.Render()
	env.ShapeX = -1
	env.Render()
}

func TestMakeAction(t *testing.T) {
	env := NewEnvironment(42)
	action := Action(MoveRight)
	env.ShapeX = 0
	if got := env.MakeAction(action); got != ok {
		t.Errorf("env.MakeAction(%v) threw error", action)
	}
	env.ShapeX = BoardWidth - (env.currentShape.getWidth() - 1)
	if got := env.MakeAction(action); got == ok {
		t.Errorf("env.MakeAction(%v) dit not threw error", action)
	}

	action = Action(MoveLeft)
	env.ShapeX = BoardWidth - (env.currentShape.getWidth() - 1)
	if got := env.MakeAction(action); got != ok {
		t.Errorf("env.MakeAction(%v) threw error", action)
	}
	action = Action(MoveLeft)
	env.ShapeX = 0
	if got := env.MakeAction(action); got == ok {
		t.Errorf("env.MakeAction(%v) dit not threw error", action)
	}
}

func (env *Environment) equal(other *Environment) bool {
	return env.Board.equal(&other.Board) &&
		env.currentShape.equal(other.currentShape) &&
		env.nextShape.equal(other.nextShape) &&
		env.ShapeX == other.ShapeX &&
		env.ShapeY == other.ShapeY &&
		env.score == other.score
}
