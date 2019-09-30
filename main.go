package main

import (
	"fmt"

	"github.com/m-nny/tetris/agents"
	"github.com/m-nny/tetris/env"
)

func main() {
	var agent env.Agent
	agent = agents.MakeLazyAgent(42)
	score, timestamp := env.Evaluate(&agent, 42)
	fmt.Printf("%v %v", score, timestamp)
}
