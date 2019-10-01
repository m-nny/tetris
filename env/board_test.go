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

func TestSquash(t *testing.T) {
	emptyBoard := Board{}
	if got, _ := emptyBoard.squash(); !got.equal(&emptyBoard) {
		t.Errorf("emptyBoard.squash() returned modified board")
	}
	original := Board{
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
		{0, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{0, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}
	wantBoard, wantScore := Board{
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
		{0, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		{0, 2, 3, 4, 5, 6, 7, 8, 9, 0},
	}, 2
	if got, score := original.squash(); !got.equal(&wantBoard) || score != wantScore {
		t.Errorf("testBoard.squash() = %v, %v, but wanted %v, %v", got, score, wantBoard, wantScore)
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
