package pin

import (
	"math/rand"
	"time"
)

// RandomPin - struct for already generated pins
type RandomPin struct {
	Generated map[string]bool
}

// Init - setup math/rand and Generated
func Init() RandomPin {
	rand.Seed(time.Now().UnixNano())
	return RandomPin{
		Generated: make(map[string]bool),
	}
}

func (r *RandomPin) GetGenerated() []string {
	keys := make([]string, 0, len(r.Generated))
	for k := range r.Generated {
		keys = append(keys, k)
	}
	return keys
}

func (r *RandomPin) GenerateOne() string {
	pin := sliceToString(r.random([]int{}, 0))
	_, ok := r.Generated[pin]
	for ok {
		pin = sliceToString(r.random([]int{}, 0))
		_, ok = r.Generated[pin]
	}
	r.Generated[pin] = true
	return pin
}

func (r *RandomPin) GenerateMultiple(number int) []string {
	list := make([]string, number)
	for i := 0; i < number; i++ {
		list[i] = r.GenerateOne()
	}
	return list
}

func (r *RandomPin) random(current []int, increment int) []int {
	random := rand.Intn(9)

	if len(current) == 0 {
		return r.random([]int{random}, 0)
	}

	prev := current[len(current)-1]

	switch {
	case len(current) == 4:
		return current
	case prev == random:
		return r.random(current, increment)
	case prev+1 == random && increment > 2:
		return r.random(current, increment)
	case prev+1 == random:
		return r.random(append(current, random), increment+1)
	default:
		return r.random(append(current, random), 0)
	}
}
