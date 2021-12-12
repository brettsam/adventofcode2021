package main

import (
	"fmt"
	"math"

	"github.com/brettsam/adventofcode2021/day5/input"
	"github.com/brettsam/adventofcode2021/day5/point"
)

func main() {
	numbers := input.GetInput()

	//fmt.Println(partOne(numbers))
	fmt.Println(partTwo(numbers))
}

func partOne(numbers []point.Vertex) int {
	var grid [1000][1000]int
	for i := 0; i < len(numbers); i++ {
		vertex := numbers[i]
		var p1x, p1y, p2x, p2y int

		// Get the lowest in p1
		if vertex.Point1.X < vertex.Point2.X {
			p1x = vertex.Point1.X
			p2x = vertex.Point2.X
		} else {
			p1x = vertex.Point2.X
			p2x = vertex.Point1.X
		}
		if vertex.Point1.Y < vertex.Point2.Y {
			p1y = vertex.Point1.Y
			p2y = vertex.Point2.Y
		} else {
			p1y = vertex.Point2.Y
			p2y = vertex.Point1.Y
		}

		if p1x == p2x {
			for y := p1y; y <= p2y; y++ {
				grid[p1x][y]++
			}
		}

		if p1y == p2y {
			for x := p1x; x <= p2x; x++ {
				grid[x][p1y]++
			}
		}
	}

	count := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid[x][y] > 1 {
				count++
			}
		}
	}

	return count
}

func partTwo(numbers []point.Vertex) int {
	var grid [1000][1000]int
	for i := 0; i < len(numbers); i++ {
		vertex := numbers[i]

		p1x := vertex.Point1.X
		p1y := vertex.Point1.Y
		p2x := vertex.Point2.X
		p2y := vertex.Point2.Y

		if p1x == p2x {
			if p1y < p2y {
				for y := p1y; y <= p2y; y++ {
					grid[p1x][y]++
				}
			} else {
				for y := p1y; y >= p2y; y-- {
					grid[p1x][y]++
				}
			}
		} else if p1y == p2y {
			if p1x < p2x {
				for x := p1x; x <= p2x; x++ {
					grid[x][p1y]++
				}
			} else {
				for x := p1x; x >= p2x; x-- {
					grid[x][p1y]++
				}
			}
		} else if math.Abs(float64(p2x-p1x)) == math.Abs(float64(p2y-p1y)) {
			if p1x < p2x && p1y < p2y {
				diff := p2x - p1x
				for x := 0; x <= diff; x++ {
					grid[p1x+x][p1y+x]++
				}
			} else if p1x < p2x && p1y > p2y {
				diff := p2x - p1x
				for x := 0; x <= diff; x++ {
					grid[p1x+x][p1y-x]++
				}
			} else if p1x > p2x && p1y < p2y {
				diff := p1x - p2x
				for x := 0; x <= diff; x++ {
					grid[p1x-x][p1y+x]++
				}
			} else if p1x > p2x && p1y > p2y {
				diff := p1x - p2x
				for x := 0; x <= diff; x++ {
					grid[p1x-x][p1y-x]++
				}
			}
		}
	}
	count := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid[x][y] > 1 {
				count++
			}
		}
	}

	return count
}
