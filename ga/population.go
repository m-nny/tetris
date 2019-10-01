package ga

import (
	"fmt"
	"math/rand"

	"github.com/m-nny/tetris/env"
)

type Population []GAgent

func InitPopulation(seed int64, size int) (pop Population) {
	rand := env.GetRand(seed)
	pop = make(Population, size)
	for i := range pop {
		subSeed := rand.Int63()
		pop[i] = *newGAgent(subSeed)
	}
	return pop
}

func (p Population) Next(seed int64, verbose bool) (newP Population) {
	N := len(p)
	rand := env.GetRand(seed)
	newP, fitness, totalF, bestF, bestA := make(Population, N), make([]float32, N), float32(0), float32(0), p[0]
	for i := 0; i < N; i++ {
		subSeed := rand.Int63()
		fitness[i] = p[i].score(subSeed)
		totalF += fitness[i]
		if fitness[i] > bestF {
			bestF = fitness[i]
			bestA = p[i]
		}
	}
	if verbose {
		fmt.Printf("Best score in population: %v by %v\n", bestF, bestA)
	}
	for i := range fitness {
		fitness[i] /= totalF
	}
	for i := 0; i < N; i++ {
		newP[i] = p.pickOne(fitness, rand)
	}
	return newP
}

func (p Population) pickOne(fitness []float32, rand *rand.Rand) GAgent {
	idx := 0
	r := rand.Float32()
	for ; r > 0; idx++ {
		r -= fitness[idx]
	}
	idx--

	return p[idx].mutate(rand)
}
