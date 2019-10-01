package env

import "fmt"

// Agent common interface
type Agent interface {
	Think(env Environment) Action
}

// Evaluate agent
func Evaluate(agent *Agent, seed int64, render bool) (score, timestamp int) {
	env := NewEnvironment(seed)
	if render {
		env.Render()
	}
	for {
		action := (*agent).Think(env)
		if err := env.MakeAction(action); err != nil {
			if render {
				fmt.Println(err)
			}
		}
		if err := env.update(); err != nil {
			break
		}
		if render {
			env.Render()
		}
	}
	return env.score, env.timestamp
}

// MultiEvaluate evaluate continously agent on multiple seeds
func MultiEvaluate(agent *Agent, seeds []int64) (scores, timestamps []int) {
	N := len(seeds)
	scores, timestamps = make([]int, N), make([]int, N)
	for i, seed := range seeds {
		scores[i], timestamps[i] = Evaluate(agent, seed, false)
	}
	return scores, timestamps
}
