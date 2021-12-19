package main

import (
	"fmt"

	"github.com/brettsam/adventofcode2021/day11/input"
	"github.com/brettsam/adventofcode2021/day11/octo"
)

var flashes int

func main() {
	numbers := input.GetInput()
	fmt.Println(partOne(numbers))

	numbers = input.GetInput()
	fmt.Println(partTwo(numbers))
}

func partOne(numbers [][]octo.Octo) int {
	for c := 1; c <= 100; c++ {
		for _, a := range numbers {
			for i := range a {
				a[i].Value += 1
			}
		}

		for y, a := range numbers {
			for x, b := range a {
				if b.Value > 9 && !b.Flashed {
					flash(numbers, x, y)
				}
			}
		}

		for y, a := range numbers {
			for x, b := range a {
				if b.Flashed {
					numbers[y][x].Value = 0
					numbers[y][x].Flashed = false
				}
			}
		}

		for _, a := range numbers {
			for _, b := range a {
				fmt.Print(b.Value)
			}
			fmt.Println()
		}
		fmt.Println()
	}

	return flashes
}

func partTwo(numbers [][]octo.Octo) int {
	i := 0
	for flashes != 100 {
		i++
		flashes = 0
		for _, a := range numbers {
			for i := range a {
				a[i].Value += 1
			}
		}

		for y, a := range numbers {
			for x, b := range a {
				if b.Value > 9 && !b.Flashed {
					flash(numbers, x, y)
				}
			}
		}

		for y, a := range numbers {
			for x, b := range a {
				if b.Flashed {
					numbers[y][x].Value = 0
					numbers[y][x].Flashed = false
				}
			}
		}

		for _, a := range numbers {
			for _, b := range a {
				fmt.Print(b.Value)
			}
			fmt.Println()
		}
		fmt.Println()
	}
	return i
}

func flash(numbers [][]octo.Octo, x int, y int) {
	numbers[y][x].Flashed = true
	flashes++
	rowLength := len(numbers[0])
	colLength := len(numbers)
	for tx := x - 1; tx <= x+1; tx++ {
		for ty := y - 1; ty <= y+1; ty++ {
			if !(tx == x && ty == y) &&
				tx >= 0 && tx <= rowLength-1 &&
				ty >= 0 && ty <= colLength-1 {
				numbers[ty][tx].Value += 1
				if numbers[ty][tx].Value > 9 && !numbers[ty][tx].Flashed {
					flash(numbers, tx, ty)
				}
			}
		}
	}
}
