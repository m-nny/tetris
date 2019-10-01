package ga

import (
	"fmt"
	"math/rand"

	"github.com/m-nny/tetris/env"
)

type GAgent struct {
	core float32
}

func newGAgent(seed int64) *GAgent {
	rand := env.GetRand(seed)
	return &GAgent{core: rand.Float32()}
}

func (gAgent *GAgent) Think(e env.Environment) env.Action {
	action := int(float32((e.ShapeX+1)*(e.ShapeY+1)) * gAgent.core * 3)
	switch action % 3 {
	case 0:
		return env.Action(env.MoveLeft)
	case 1:
		return env.Action(env.MoveRight)
	case 2:
		return env.Action(env.DontMove)
	}
	return env.Action(env.DontMove)
}

func (gAgent *GAgent) score(seed int64) float32 {
	var agent env.Agent
	agent = gAgent

	score, lt := env.Evaluate(&agent, seed, false)
	fitness := float32(score*score*score) + float32(lt)
	return fitness

}

func (gAgent *GAgent) mutate(rand *rand.Rand) (baby GAgent) {
	baby = *gAgent
	if r := rand.Float32(); r <= mutationRate {
		dx := float32(rand.NormFloat64()) / 2
		baby.core += dx
		fmt.Printf("baby mutated: %v + %v = %v\n", gAgent.core, dx, baby.core)
	}
	return baby
}
