package evoli

import (
	"math/rand"
)

type proportionalToRankSelecter struct{}

func (s proportionalToRankSelecter) Select(pop Population, survivorsSize int) (Population, error) {
	err := checkParams(pop, survivorsSize)
	if err != nil {
		return nil, err
	}
	if survivorsSize >= pop.Len() {
		return pop, nil
	}
	newPop := pop.New(pop.Cap())
	totalScore := s.computeTotalScore(pop)
	pop.Sort()
	for newPop.Len() < survivorsSize {
		for i := 0; i < pop.Len(); i++ {
			if newPop.Len() >= survivorsSize {
				break
			}
			score := float64(pop.Len() - i)
			if rand.Float64() <= score/totalScore {
				indiv, _ := pop.Get(i)
				pop.RemoveAt(i)
				newPop.Add(indiv)
				totalScore -= score
			}
		}
	}
	return newPop, nil
}

func (s proportionalToRankSelecter) computeTotalScore(pop Population) float64 {
	n := float64(pop.Len())
	return n * (n + 1) / 2 // 1+2+3+...+n
}

// NewProportionalToRankSelecter is the constructor for selecter based on ranking across the population
func NewProportionalToRankSelecter() Selecter {
	return proportionalToRankSelecter{}
}
