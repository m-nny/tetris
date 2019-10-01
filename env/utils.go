package env

import (
	"math/rand"
)

func GetRand(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}
