package main

import (
	"fmt"
	"math"

	"github.com/brettsam/adventofcode2021/day7/input"
)

func main() {
	numbers := input.GetInput()

	fmt.Println(partOne(numbers))
	fmt.Println(partTwo(numbers))
}

func partOne(numbers []int) int {
	least := math.MaxInt32
	for i := 0; i < len(numbers); i++ {
		current := 0
		for j := 0; j < len(numbers); j++ {
			// try to move the number at j to i
			current += int(math.Abs(float64(numbers[j] - i)))
		}
		if current < least {
			least = current
		}
	}
	return least
}

func partTwo(numbers []int) int {
	least := math.MaxInt32
	for i := 0; i < len(numbers); i++ {
		current := 0
		for j := 0; j < len(numbers); j++ {
			// try to move the number at j to i
			distance := int(math.Abs(float64(numbers[j] - i)))
			for k := 1; k <= distance; k++ {
				current += k
			}
		}
		if current < least {
			least = current
		}
	}
	return least
}
