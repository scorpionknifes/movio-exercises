package main

import (
	"fmt"

	"github.com/scorpionknifes/movio/merge"
	"github.com/scorpionknifes/movio/pin"
)

func main() {
	m := merge.SortedMerge([]int{1, 1, 3, 5}, []int{1, 2, 3, 4})
	fmt.Println(m)

	p := pin.Init()
	p.GenerateOneThousand()
}
