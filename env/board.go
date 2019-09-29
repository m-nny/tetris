package env

const (
	boardWidth  = 10
	boardHeight = 15
)

// Board current game
type Board [boardHeight][boardWidth]int

func (board *Board) canFit(shape Shape, x1, y1 int) error {
	x2, y2 := x1+shape.getWidth()-1, y1+shape.getHeight()-1
	if !(0 <= x1 && x1 < boardWidth && 0 <= y1 && y1 < boardHeight) ||
		!(0 <= x2 && x2 < boardWidth && 0 <= y2 && y2 < boardHeight) {
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
