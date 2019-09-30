package env

// Action ...
type Action int

// DontMove ...
const (
	DontMove  Action = iota
	MoveLeft         = iota
	MoveRight        = iota
	MoveDown         = iota
)
