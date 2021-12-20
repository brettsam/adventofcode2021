package main

import (
	"fmt"
	"math"

	"github.com/brettsam/adventofcode2021/day14/input"
	"github.com/brettsam/adventofcode2021/day14/polymer"
)

func main() {
	numbers := input.GetInput()
	fmt.Println(partOne(numbers, "PPFCHPFNCKOKOSBVCFPP"))

	numbers = input.GetInput()
	fmt.Println(partTwo(numbers, "PPFCHPFNCKOKOSBVCFPP"))
}

func partOne(numbers map[string]string, input string) uint64 {
	polymer := polymer.Polymer{Inserts: numbers}
	counts := polymer.Generate(input, 10)

	return doMath(counts)
}

func partTwo(numbers map[string]string, input string) uint64 {
	polymer := polymer.Polymer{Inserts: numbers}
	counts := polymer.Generate(input, 40)

	return doMath(counts)
}

func doMath(m map[string]uint64) uint64 {
	low := uint64(math.MaxUint64)
	high := uint64(0)

	for _, value := range m {
		if value < low {
			low = value
		}
		if value > high {
			high = value
		}
	}
	return high - low
}
