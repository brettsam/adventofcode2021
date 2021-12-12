package main

import (
	"fmt"

	"github.com/brettsam/adventofcode2021/day6/input"
)

func main() {
	numbers := input.GetInput()

	fmt.Println(partOne(numbers))
	fmt.Println(partTwo(numbers))
}

func partOne(numbers []int) int {
	for day := 0; day < 80; day++ {
		numbers = simulateDay(numbers)
	}

	return len(numbers)
}

func partTwo(numbers []int) int {
	var counts [9]int
	for i := 0; i < len(numbers); i++ {
		counts[numbers[i]]++
	}
	for day := 0; day < 256; day++ {
		zero := counts[0]
		one := counts[1]
		two := counts[2]
		three := counts[3]
		four := counts[4]
		five := counts[5]
		six := counts[6]
		seven := counts[7]
		eight := counts[8]

		counts[0] = one
		counts[1] = two
		counts[2] = three
		counts[3] = four
		counts[4] = five
		counts[5] = six
		counts[6] = seven + zero
		counts[7] = eight
		counts[8] = zero
	}

	sum := 0
	for c := 0; c < 9; c++ {
		sum += counts[c]
	}

	return sum
}

func simulateDay(numbers []int) []int {
	length := len(numbers)
	for i := 0; i < length; i++ {
		if numbers[i] == 0 {
			numbers[i] = 6
			numbers = append(numbers, 8)
		} else {
			numbers[i]--
		}
	}
	return numbers
}
