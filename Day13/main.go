package main

import (
	"fmt"

	"github.com/brettsam/adventofcode2021/day13/input"
)

func main() {
	numbers := input.GetInput()
	fmt.Println(partOne(numbers))

	numbers = input.GetInput()
	partTwo(numbers)
}

func partOne(numbers [][]string) int {
	numbers = foldLeft(numbers, 655)

	return countDots(numbers)
}

func partTwo(numbers [][]string) {
	/*
		fold along x=655
		fold along y=447
		fold along x=327
		fold along y=223
		fold along x=163
		fold along y=111
		fold along x=81
		fold along y=55
		fold along x=40
		fold along y=27
		fold along y=13
		fold along y=6
	*/
	numbers = foldLeft(numbers, 655)
	numbers = foldUp(numbers, 447)
	numbers = foldLeft(numbers, 327)
	numbers = foldUp(numbers, 223)
	numbers = foldLeft(numbers, 163)
	numbers = foldUp(numbers, 111)
	numbers = foldLeft(numbers, 81)
	numbers = foldUp(numbers, 55)
	numbers = foldLeft(numbers, 40)
	numbers = foldUp(numbers, 27)
	numbers = foldUp(numbers, 13)
	numbers = foldUp(numbers, 6)

	printBoard(numbers)
}

func printBoard(board [][]string) {
	for _, a := range board {
		for _, b := range a {
			fmt.Print(b)
		}
		fmt.Println()
	}
	fmt.Println()
}

func foldUp(numbers [][]string, index int) [][]string {
	// copy everything up to index
	newBoard := make([][]string, index)
	for y := 0; y < index; y++ {
		newBoard[y] = numbers[y]
	}

	for y := 1; y <= index; y++ {
		for x := range numbers[y] {
			c := numbers[index+y][x]
			if c == "#" {
				newBoard[index-y][x] = "#"
			}
		}
	}

	return newBoard
}

func countDots(numbers [][]string) int {
	count := 0

	for _, a := range numbers {
		for _, b := range a {
			if b == "#" {
				count++
			}
		}
	}

	return count
}

func foldLeft(numbers [][]string, index int) [][]string {
	// copy everything up to index
	l := len(numbers)
	newBoard := make([][]string, l)
	for y := 0; y < l; y++ {
		newBoard[y] = numbers[y][0:index]
	}

	// now copy everything from the right over
	for y := 0; y < l; y++ {
		rightSide := numbers[y][index:]
		rsl := len(rightSide) - 1
		for x := 0; x < index; x++ {
			if newBoard[y][x] != "#" {
				ri := rsl - x
				if ri >= 0 {
					newBoard[y][x] = rightSide[ri]
				}
			}
		}
	}

	return newBoard
}
