package main

import (
	"fmt"

	"github.com/brettsam/adventofcode2021/day4/input"
)

func main() {
	numbers, boards := input.GetInput()

	fmt.Println(numbers)

	copy1 := make([][5][5]int, len(boards))
	copy2 := make([][5][5]int, len(boards))

	copy(copy1, boards)
	copy(copy2, boards)

	fmt.Println(partOne(numbers, copy1))
	fmt.Println(partTwo(numbers, copy2))
}

func partOne(numbers []int, boards [][5][5]int) int {
	for i := 0; i < len(numbers); i++ {
		var number = numbers[i]
		for j := 0; j < len(boards); j++ {
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					if boards[j][k][l] == number {
						boards[j][k][l] = -1
					}
				}
			}
		}

		getWinner := getWinner(boards)
		if getWinner > 0 {
			fmt.Println(number, " ", getWinner)
			return getWinner * number
		}
	}

	return 0
}

func partTwo(numbers []int, boards [][5][5]int) int {
	winnerValue := 0
	winners := make([]bool, len(boards))
	for i := 0; i < len(numbers); i++ {
		var number = numbers[i]
		for j := 0; j < len(boards); j++ {
			if winners[j] {
				continue
			}

			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					if boards[j][k][l] == number {
						boards[j][k][l] = -1
					}
				}
			}

			winnerSum := getWinnerSum(boards[j])
			if winnerSum > 0 {
				winnerValue = winnerSum * number
				winners[j] = true
			}
		}
	}
	return winnerValue
}

func getWinner(boards [][5][5]int) int {
	for j := 0; j < len(boards); j++ {
		boardWinner := getWinnerSum(boards[j])
		if boardWinner > 0 {
			return boardWinner
		}
	}
	return 0
}

func getWinnerSum(board [5][5]int) int {
	boardWinner := false
	sum := 0

	for k := 0; k < 5; k++ {
		rowWinner := true
		columnWinner := true
		for l := 0; l < 5; l++ {
			if board[k][l] > 0 {
				sum += board[k][l]
				rowWinner = false
			}

			if board[l][k] > 0 {
				columnWinner = false
			}
		}

		if rowWinner || columnWinner {
			boardWinner = true
		}
	}

	if boardWinner {
		return sum
	}

	return 0
}
