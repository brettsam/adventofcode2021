package main

import (
	"fmt"

	"github.com/brettsam/adventofcode2021/day15/input"
)

func main() {
	start := input.GetInput()
	fmt.Println(partOne(start))

	start = input.GetInput()
	fmt.Println(partTwo(start))
}

func partOne(cave [][]*input.Node) int {
	curX := len(cave[0]) - 1
	curY := len(cave) - 1

	exit := cave[curY][curX]
	exit.Sum = exit.Risk
	list := []*input.Node{exit}

	for {
		for _, cell := range list {
			list = evaluateCell(cave, cell.X, cell.Y-1, cell.Sum, list)
			list = evaluateCell(cave, cell.X-1, cell.Y, cell.Sum, list)
			list = evaluateCell(cave, cell.X, cell.Y+1, cell.Sum, list)
			list = evaluateCell(cave, cell.X+1, cell.Y, cell.Sum, list)

			list = list[1:]
		}

		if len(list) == 0 {
			break
		}
	}

	entrance := cave[0][0]
	entrance.Sum -= entrance.Risk
	return entrance.Sum
}

func partTwo(cave [][]*input.Node) int {
	lenX := len(cave[0])
	lenY := len(cave)

	for _, row := range cave {
		for y := 1; y <= 4; y++ {
			newRow := make([]*input.Node, 0)
			for _, cell := range row {
				risk := cell.Risk + y
				if risk > 9 {
					risk -= 9
				}
				offsetY := y * lenY
				newRow = append(newRow, &input.Node{Risk: risk, X: cell.X, Y: cell.Y + offsetY})
			}
			cave = append(cave, newRow)
		}
	}

	for y, row := range cave {
		for x := 1; x <= 4; x++ {
			for _, cell := range row {
				risk := cell.Risk + x
				if risk > 9 {
					risk -= 9
				}
				offsetX := x * lenX
				cave[y] = append(cave[y], &input.Node{Risk: risk, X: cell.X + offsetX, Y: cell.Y})
			}
		}
	}

	for _, row := range cave {
		for _, cell := range row {
			fmt.Print(cell.Risk)
		}
		fmt.Println()
	}

	return partOne(cave)
}

func evaluateCell(cave [][]*input.Node, x int, y int, curSum int, list []*input.Node) []*input.Node {
	if x < 0 || y < 0 || x > len(cave[0])-1 || y > len(cave)-1 {
		return list
	}

	cell := cave[y][x]
	newSum := cell.Risk + curSum
	if cell.Sum == 0 || newSum < cell.Sum {
		cell.Sum = newSum
		return append(list, cell)
	}
	return list
}
