package env

import "math/rand"

// Shape ...
type Shape [][]int

var (
	emptyShape = fill(Shape{{0}}, 0)
	iShape     = fill(Shape{{1, 1, 1, 1}}, 1)
	oShape     = fill(Shape{{1, 1}, {1, 1}}, 2)
	tShape     = fill(Shape{{1, 1, 1}, {0, 1, 0}}, 3)
	jShape     = fill(Shape{{1, 1, 1}, {0, 0, 1}}, 4)
	lShape     = fill(Shape{{1, 1, 1}, {1, 0, 0}}, 5)
	sShape     = fill(Shape{{0, 1, 1}, {1, 1, 0}}, 6)
	zShape     = fill(Shape{{1, 1, 0}, {0, 1, 1}}, 7)
)
var playableShapes = [...]Shape{iShape, oShape, tShape, jShape, lShape, sShape, zShape}

func fill(shape Shape, scale int) Shape {
	for i, row := range shape {
		for j := range row {
			shape[i][j] *= scale
		}
	}
	return shape
}

func (shape Shape) getHeight() int {
	return len(shape)
}

func (shape Shape) getWidth() int {
	return len(shape[0])
}

func getRandomShape() Shape {
	idx := rand.Intn(len(playableShapes))
	return playableShapes[idx]
}
