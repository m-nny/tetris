package env

import "testing"

func TestCanFit(t *testing.T) {
	board := Board{}
	ok, shape, x, y := error(nil), oShape, 0, 0
	if got := board.canFit(shape, x, y); got != ok {
		t.Errorf("emptyBoard.canFit(%v, %v, %v) threw error %v", shape, x, y, got)
	}
	board[x][y] = 1
	if got := board.canFit(shape, x, y); got == ok {
		t.Errorf("emptyBoard.canFit(%v, %v, %v) did not threw error", shape, x, y)
	}

	x, y = boardWidth, boardHeight
	if got := board.canFit(shape, x, y); got == ok {
		t.Errorf("emptyBoard.canFit(%v, %v, %v) did not threw error", shape, x, y)
	}
}

func TestFit(t *testing.T) {
	board := Board{}
	want, shape, x, y := error(nil), oShape, 0, 0
	if _, got := board.fit(shape, x, y); got != want {
		t.Errorf("emptyBoard.fit(%v, %v, %v) = %v, want %v", shape, x, y, got, want)
	}
	if _, got := board.fit(shape, x, y); got != want {
		t.Errorf("emptyBoard.fit() should not mutate board")
	}

	board[x][y] = 1

	dontWant := error(nil)
	if _, got := board.fit(shape, x, y); got == dontWant {
		t.Errorf("emptyBoard.fit(%v, %v, %v) = %v, but it should not be", shape, x, y, got)
	}
}

func (board *Board) equal(other *Board) bool {
	for i, row := range board {
		for j := range row {
			if board[i][j] != other[i][j] {
				return false
			}
		}
	}
	return true
}
