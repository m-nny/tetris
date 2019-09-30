package env

import "fmt"

// Agent common interface
type Agent interface {
	Think(env Environment) Action
}

// Evaluate agent
func Evaluate(agent *Agent, seed int64) (score, timestamp int) {
	env := NewEnvironment(seed)
	for {
		action := (*agent).Think(env)
		if err := env.MakeAction(action); err != nil {
			fmt.Println(err)
		}
		if err := env.Update(); err != nil {
			break
		}
		env.Render()
	}
	return env.score, env.timestamp
}
