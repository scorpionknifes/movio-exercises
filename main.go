package main

import (
	"fmt"

	"github.com/scorpionknifes/movio/merge"
	"github.com/scorpionknifes/movio/pin"
)

func main() {
	// exercise1
	m := merge.SortedMerge([]int{1, 1, 3, 5}, []int{1, 2, 3, 4})
	fmt.Println(m)

	// exercise2
	p := pin.Init()
	c := p.GenerateOneThousand()
	fmt.Println(c)
}
