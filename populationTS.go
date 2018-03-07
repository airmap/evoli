package evoli

import "sync"

type populationTS struct {
	population
	mutex sync.RWMutex
}

// NewPopulationTS creates a threadsafe population
func NewPopulationTS(capacity int) Population {
	pop := NewPopulation(capacity)
	return &populationTS{
		*pop.(*population),
		sync.RWMutex{},
	}
}

// Len returns the current livings count of a population
func (p *populationTS) Len() int {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.population.Len()
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p *populationTS) Less(i, j int) bool {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.population.Less(i, j)
}

// Swap swaps the elements with indexes i and j.
func (p *populationTS) Swap(i, j int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.population.Swap(i, j)
}

// Sort sort the population
func (p *populationTS) Sort() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.population.Sort()
}

// SetCap set the resize the population capacity
func (p *populationTS) SetCap(newCap int) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.population.SetCap(newCap)
}

// Add adds an individual to a population. If the populagtion has already reached its capacity, capacity is incremented.
func (p *populationTS) Add(indiv ...Individual) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.population.Add(indiv...)
}

// Get returns the individual at index i
func (p *populationTS) Get(i int) (Individual, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.population.Get(i)
}

// RemoveAt removes and returns the individual at index i
func (p *populationTS) RemoveAt(i int) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.population.RemoveAt(i)
}

// Remove removes all given individuals
func (p *populationTS) Remove(individuals ...Individual) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.population.Remove(individuals...)
}

// Replace replaces and returns the individual at index i by the substitute
func (p *populationTS) Replace(i int, substitute Individual) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.population.Replace(i, substitute)
}

// Min returns the least Resilent individual
func (p *populationTS) Min() Individual {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.population.Min()
}

// Max returns the most Resilent individual
func (p *populationTS) Max() Individual {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.population.Max()
}

// Has return true if the specified individual is in the population
func (p *populationTS) Has(individuals ...Individual) bool {
	has := true
	for _, indiv := range individuals {
		_, err := p.IndexOf(indiv)
		has = has && err == nil
	}
	return has
}

// IndexOf returns the inde of the specified individual if it exists
func (p *populationTS) IndexOf(indiv Individual) (int, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.population.IndexOf(indiv)
}

// Each traverse the population and execute given callback on each individual. Stops if the callbak return false.
func (p *populationTS) Each(f func(indiv Individual) bool) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	p.population.Each(f)
}

// Slice returns the population as []Individual
func (p *populationTS) Slice() []Individual {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.population.Slice()
}

func (p *populationTS) New(cap int) Population {
	return NewPopulationTS(cap)
}
