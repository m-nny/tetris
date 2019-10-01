package main

import (
	"fmt"
	"math/rand"

	"github.com/m-nny/tetris/agents"
	"github.com/m-nny/tetris/env"
	"github.com/m-nny/tetris/ga"
)

func meanScore(agent *env.Agent, n int) (mScore, mLT float32) {
	for i := 0; i < n; i++ {
		seed := rand.Int63()
		score, lt := env.Evaluate(agent, seed, false)
		mScore += float32(score)
		mLT += float32(lt)
	}
	mScore /= float32(n)
	mLT /= float32(n)
	fmt.Printf("mean over %v score: %v lifetime: %v\n", n, mScore, mLT)
	return mScore, mLT
}

func testRandomAgent() {
	var agent env.Agent
	seeds := []int64{42, 41, 40}

	agent = agents.MakeRandomAgent(seeds[0])
	scores, lf := env.MultiEvaluate(&agent, seeds[:])
	for i := range seeds {
		fmt.Printf("seed: %v score: %v lifetime: %v\n", seeds[i], scores[i], lf[i])
	}

	N := 100
	meanScore(&agent, N)
}

// func testGAgent() {
// 	var agent env.Agent
// 	seed := int64(42)
// 	agent = ga.NewGAgent(seed)
// 	// score, lf := env.Evaluate(&agent, seed, true)
// 	// fmt.Printf("seed: %v score: %v lifetime: %v\n", seed, score, lf)
// 	N := 100
// 	meanScore(&agent, N)
// }

func testGA() {
	p := ga.InitPopulation(42, 1000)
	// fmt.Printf("%v\n", p)
	newP := p.Next(40, true)
	fmt.Printf("%v\n", newP)
}

func main() {
	testGA()
}
