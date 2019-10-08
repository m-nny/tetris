package env

const (
	BoardWidth  = 10
	BoardHeight = 15
)

// Board current game
type Board [BoardHeight][BoardWidth]int

func (board *Board) canFit(shape Shape, x1, y1 int) error {
	x2, y2 := x1+shape.getWidth()-1, y1+shape.getHeight()-1
	if !(0 <= x1 && x1 < BoardWidth && 0 <= y1 && y1 < BoardHeight) ||
		!(0 <= x2 && x2 < BoardWidth && 0 <= y2 && y2 < BoardHeight) {
		return errBounds()
	}
	for x, dx := x1, 0; x <= x2; x, dx = x+1, dx+1 {
		for y, dy := y1, 0; y <= y2; y, dy = y+1, dy+1 {
			if shape[dy][dx] > 0 && board[y][x] > 0 {
				return errFilled()
			}
		}
	}
	return nil
}

func (board *Board) fit(shape Shape, x1, y1 int) (*Board, error) {
	if err := board.canFit(shape, x1, y1); err != nil {
		return nil, err
	}
	x2, y2 := x1+shape.getWidth()-1, y1+shape.getHeight()-1
	newBoard := *board
	for x, dx := x1, 0; x <= x2; x, dx = x+1, dx+1 {
		for y, dy := y1, 0; y <= y2; y, dy = y+1, dy+1 {
			if shape[dy][dx] > 0 {
				newBoard[y][x] = shape[dy][dx]
			}
		}
	}
	return &newBoard, nil
}

func (board *Board) squash() (newBoard *Board, squash int) {
	var allFilled [BoardHeight]bool
	for y := 0; y < BoardHeight; y++ {
		allFilled[y] = true
		for x := 0; x < BoardWidth; x++ {
			if board[y][x] == 0 {
				allFilled[y] = false
			}
		}
		if allFilled[y] {
			squash++
		}
	}
	newBoard = &Board{}
	yTarget := BoardHeight - 1
	for ySource := BoardHeight - 1; ySource >= 0; ySource-- {
		if allFilled[ySource] {
			continue
		}
		for x := 0; x < BoardWidth; x++ {
			newBoard[yTarget][x] = board[ySource][x]
		}
		yTarget--
	}
	return newBoard, squash
}
