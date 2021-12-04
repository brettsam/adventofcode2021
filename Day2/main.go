package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brettsam/adventofcode2021/day2/input"
)

func main() {
	measurements := input.GetDirections()

	fmt.Println(partOne(measurements))
	fmt.Println(partTwo(measurements))
}

func partOne(measurements []string) int {
	x := 0
	y := 0
	for i := 0; i < len(measurements); i++ {
		instruction := measurements[i]

		split := strings.Split(instruction, " ")
		i, _ := strconv.ParseInt(split[1], 10, 32)
		amount := int(i)

		switch split[0] {

		case "forward":
			x += amount
		case "down":
			y += amount
		case "up":
			y -= amount
		}
	}
	return x * y
}

func partTwo(measurements []string) int {
	x := 0
	y := 0
	a := 0
	for i := 0; i < len(measurements); i++ {
		instruction := measurements[i]

		split := strings.Split(instruction, " ")
		i, _ := strconv.ParseInt(split[1], 10, 32)
		amount := int(i)

		switch split[0] {

		case "forward":
			x += amount
			y += amount * a
		case "down":
			a += amount
		case "up":
			a -= amount
		}
	}
	return x * y
}
