package agents

import (
	"math/rand"

	"github.com/m-nny/tetris/env"
)

type RandomAgent struct {
	rand *rand.Rand
}

func MakeLazyAgent(seed int64) *RandomAgent {
	r := rand.NewSource(seed)
	return &RandomAgent{rand.New(r)}
}

func (agent *RandomAgent) Think(e env.Environment) env.Action {
	actionIdx := agent.rand.Intn(4)
	return env.Action(actionIdx)
}
