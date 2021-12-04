package main

import (
	"fmt"

	"github.com/brettsam/adventofcode2021/Day1/input"
)

func main() {
	measurements := input.GetMeasurements()

	fmt.Println(partOne(measurements))
	fmt.Println(partTwo(measurements))
}

func partOne(measurements []int) int {
	count := 0
	for i := 1; i < len(measurements); i++ {
		last := measurements[i-1]
		curr := measurements[i]
		if curr > last {
			count++
		}
	}
	return count
}

func partTwo(measurements []int) int {
	count := 0
	for i := 3; i < len(measurements); i++ {
		one := measurements[i-3]
		two := measurements[i-2]
		three := measurements[i-1]
		four := measurements[i]
		if four+three+two > three+two+one {
			count++
		}
	}
	return count
}
