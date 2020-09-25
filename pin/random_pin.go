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

// GetGenerated - used to get all generated pins by generator
func (r *RandomPin) GetGenerated() []string {
	keys := make([]string, 0, len(r.Generated))
	for k := range r.Generated {
		keys = append(keys, k)
	}
	return keys
}

// GenerateOne - generate one pin
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

// GenerateMultiple - generate n numbers of pins
func (r *RandomPin) GenerateMultiple(number int) []string {
	list := make([]string, number)
	for i := 0; i < number; i++ {
		list[i] = r.GenerateOne()
	}
	return list
}

// GenerateOneThousand - generate exactly 1000 pins
func (r *RandomPin) GenerateOneThousand() []string {
	return r.GenerateMultiple(1000)
}

// random - generates the next digit in a pin (recursive call)
func (r *RandomPin) random(current []int, increment int) []int {
	random := rand.Intn(9)

	if len(current) == 0 {
		return r.random([]int{random}, 0)
	}

	if len(current) == 4 {
		return current
	}

	return r.random(r.increment(current, random, increment))

}

// increment - check if follows requirement
func (r *RandomPin) increment(current []int, random int, increment int) ([]int, int) {
	prev := current[len(current)-1]

	switch {
	case prev == random:
		return current, increment
	case prev+1 == random && increment > 2:
		return current, increment
	case prev+1 == random:
		return append(current, random), increment + 1
	default:
		return append(current, random), 0
	}
}
