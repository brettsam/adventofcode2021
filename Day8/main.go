package main

import (
	"fmt"
	"strconv"

	"github.com/brettsam/adventofcode2021/day8/input"
	"github.com/brettsam/adventofcode2021/day8/line"
)

func main() {
	numbers := input.GetInput()

	fmt.Println(partOne(numbers))
	fmt.Println(partTwo(numbers))
}

func partOne(numbers []line.Line) int {
	count := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < 4; j++ {
			o := numbers[i].Output[j]
			l := len(o)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				count++
			}
		}
	}
	return count
}

func partTwo(numbers []line.Line) int {
	count := 0
	for i := 0; i < len(numbers); i++ {
		var mapping [10]string

		var top string
		var tr string
		var br string

		// possibilities
		var rightSide []string

		var tl []string
		var mid []string

		// first find the 1
		for k := 0; k < 10; k++ {
			inputString := numbers[i].Input[k]

			if len(inputString) == 2 {
				mapping[1] = inputString
				rightSide = append(rightSide, string(inputString[0]), string(inputString[1]))
			}

			if len(inputString) == 7 {
				mapping[8] = inputString
			}
		}

		// then the 7
		for k := 0; k < 10; k++ {
			inputString := numbers[i].Input[k]

			if len(inputString) == 3 {
				mapping[7] = inputString
				for v := 0; v < 3; v++ {
					if !contains(rightSide, inputString[v:v+1]) {
						top = inputString[v : v+1]
						break
					}
				}
				break
			}
		}

		// now do 4
		for k := 0; k < 10; k++ {
			inputString := numbers[i].Input[k]

			if len(inputString) == 4 {
				mapping[4] = inputString
				for v := 0; v < 4; v++ {
					if !contains(rightSide, inputString[v:v+1]) {
						tl = append(tl, inputString[v:v+1])
						mid = append(mid, inputString[v:v+1])
					}
				}
			}
		}

		// find 9
		for k := 0; k < 10; k++ {
			inputString := numbers[i].Input[k]
			if len(inputString) == 6 {
				matches := 0
				for v := 0; v < 6; v++ {
					c := inputString[v : v+1]
					if top == c || contains(rightSide, c) || contains(mid, c) {
						matches++
					}
				}

				if matches == 5 {
					mapping[9] = inputString
				}
			}
		}

		// find 6 and 0
		for k := 0; k < 10; k++ {
			inputString := numbers[i].Input[k]
			if inputString != mapping[9] && len(inputString) == 6 {
				matches := 0
				brTemp := ""
				for v := 0; v < 6; v++ {
					c := inputString[v : v+1]

					// if it contains tr and br, it's 0; otherwise 6
					if contains(rightSide, c) {
						matches++
						brTemp = c
					}
				}

				if matches == 2 {
					mapping[0] = inputString
				} else {
					mapping[6] = inputString
					br = brTemp
					if rightSide[0] == br {
						tr = rightSide[1]
					} else {
						tr = rightSide[0]
					}
				}
			}
		}

		for k := 0; k < 10; k++ {
			inputString := numbers[i].Input[k]
			if inputString != mapping[9] && len(inputString) == 5 {
				hasTr := false
				hasBr := false
				for v := 0; v < 5; v++ {
					c := inputString[v : v+1]

					// if it contains tr and br, it's 3
					if c == tr {
						hasTr = true
					}

					if c == br {
						hasBr = true
					}
				}

				if hasBr && hasTr {
					mapping[3] = inputString
				} else if hasBr {
					mapping[5] = inputString
				} else if hasTr {
					mapping[2] = inputString
				}
			}
		}

		endingNumber := ""
		for j := 0; j < 4; j++ {
			op := numbers[i].Output
			o := op[j]
			for v := 0; v < 10; v++ {
				if mapping[v] == o {
					si := strconv.Itoa(v)
					endingNumber = endingNumber + si
				}
			}
		}
		y, _ := strconv.ParseInt(endingNumber, 10, 32)
		count += int(y)
	}
	return count
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
