package main

import (
	"fmt"

	"github.com/brettsam/adventofcode2021/Day1/input"
)

func main() {
	measurements := input.GetMeasurements()
	count := 0
	for i := 1; i < len(measurements); i++ {
		last := measurements[i-1]
		curr := measurements[i]
		if curr > last {
			count++
		}
	}
	fmt.Println(count)
}
