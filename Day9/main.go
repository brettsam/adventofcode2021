package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/brettsam/adventofcode2021/day9/input"
)

func main() {
	numbers := input.GetInput()

	fmt.Println(partOne(numbers))
	fmt.Println(partTwo(numbers))
}

func partOne(numbers [][]int) int {
	count := 0
	maxX := len(numbers[0]) - 1
	maxY := len(numbers) - 1
	for y, c := range numbers {
		for x, _ := range c {
			up := math.MaxInt32
			down := math.MaxInt32
			left := math.MaxInt32
			right := math.MaxInt32

			if x >= 1 {
				left = numbers[y][x-1]
			}

			if x < maxX {
				right = numbers[y][x+1]
			}

			if y >= 1 {
				up = numbers[y-1][x]
			}

			if y < maxY {
				down = numbers[y+1][x]
			}
			cur := numbers[y][x]
			if cur < up && cur < down && cur < left && cur < right {
				count += 1 + cur
			}
		}
	}
	return count
}

func partTwo(numbers [][]int) int {
	var basins []int
	for y, c := range numbers {
		for x := range c {
			basinSize := evaluateNeighbors(numbers, x, y)
			if basinSize > 0 {
				basins = append(basins, basinSize)
			}
		}
	}
	sort.SliceStable(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})
	return basins[0] * basins[1] * basins[2]
}

func evaluateNeighbors(numbers [][]int, startX int, startY int) int {
	count := 0

	cur := numbers[startY][startX]

	if cur == 9 {
		return count
	}

	// add self to count
	if cur != -1 {
		numbers[startY][startX] = -1
		count += 1
	} else {
		return count
	}

	// go left
	if startX > 0 {
		count += evaluateNeighbors(numbers, startX-1, startY)
	}

	// go right
	if startX < len(numbers[0])-1 {
		count += evaluateNeighbors(numbers, startX+1, startY)
	}

	// go up
	if startY > 0 {
		count += evaluateNeighbors(numbers, startX, startY-1)
	}

	// go down
	if startY < len(numbers)-1 {
		count += evaluateNeighbors(numbers, startX, startY+1)
	}

	return count
}
