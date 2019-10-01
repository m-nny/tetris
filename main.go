package main

import (
	"fmt"
	"math/rand"

	"github.com/m-nny/tetris/agents"
	"github.com/m-nny/tetris/env"
)

func meanScore(agent *env.Agent, n int) (meanScore, meanLT float32) {
	for i := 0; i < n; i++ {
		seed := rand.Int63()
		score, lt := env.Evaluate(agent, seed, false)
		meanScore += float32(score)
		meanLT += float32(lt)
	}
	meanScore /= float32(n)
	meanLT /= float32(n)
	return meanScore, meanLT
}

func main() {
	var agent env.Agent
	seeds := []int64{42, 41, 40}

	agent = agents.MakeLazyAgent(seeds[0])
	scores, lf := env.MultiEvaluate(&agent, seeds[:])
	for i := range seeds {
		fmt.Printf("seed: %v score: %v lifetime: %v\n", seeds[i], scores[i], lf[i])
	}

	N := 100
	mScore, mLT := meanScore(&agent, N)
	fmt.Printf("mean over %v score: %v lifetime: %v\n", N, mScore, mLT)
}
