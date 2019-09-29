package env

import "testing"

func TestFill(t *testing.T) {
	shape, scale := Shape{{1, 0}, {0, 1}}, 10
	want := Shape{{scale, 0}, {0, scale}}
	if got := fill(shape, scale); !got.equal(want) {
		t.Errorf("fill(%v, %v) = %v, want %v", shape, scale, got, want)
	}
}

func TestGetDimentions(t *testing.T) {
	shape := Shape{{1, 0}, {0, 1}, {1, 1}, {0, 0}}
	want := 4
	if got := shape.getHeight(); got != want {
		t.Errorf("%v.getHeight() = %v, want %v", shape, got, want)
	}
	want = 2
	if got := shape.getWidth(); got != want {
		t.Errorf("%v.getWidth() = %v, want %v", shape, got, want)
	}
}

func TestGetShape(t *testing.T) {
	rShape := getRandomShape()
	got, want := false, true
	for _, pShape := range playableShapes {
		if rShape.equal(pShape) {
			got = true
			break
		}
	}
	if got != want {
		t.Errorf("getRandomShape() = %v, returned unplayable shape", rShape)
	}
}

func (shape Shape) equal(other Shape) bool {
	if shape.getHeight() != other.getHeight() {
		return false
	}
	if shape.getWidth() != other.getWidth() {
		return false
	}
	for i, row := range shape {
		for j := range row {
			if shape[i][j] != other[i][j] {
				return false
			}
		}
	}
	return true
}
