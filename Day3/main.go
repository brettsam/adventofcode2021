package main

import (
	"fmt"

	"github.com/brettsam/adventofcode2021/day3/input"
)

func main() {
	input := input.GetInput()

	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func partOne(input []string) int {
	var tracker = [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var gamma int = 0
	var epsilon int = 0

	for i := 0; i < len(input); i++ {
		chars := []rune(input[i])
		for j := 0; j < len(chars); j++ {
			if string(chars[j]) == "1" {
				tracker[j]++
			} else {
				tracker[j]--
			}
		}
	}

	for k := 0; k < len(tracker); k++ {
		var pos = 11 - k
		if tracker[k] > 0 {
			gamma = setBit(gamma, pos)
		} else {
			epsilon = setBit(epsilon, pos)
		}
	}

	return gamma * epsilon
}

func partTwo(input []string) int {
	var o int = 0
	var co2 int = 0

	most := filter(input, true, 0)
	least := filter(input, false, 0)

	for x := 1; x < 12; x++ {
		if len(most) > 1 {
			most = filter(most, true, x)
		}
		if len(least) > 1 {
			least = filter(least, false, x)
		}
	}

	o = bitStringToInt(most[0])
	co2 = bitStringToInt(least[0])

	return o * co2
}

func filter(input []string, useMost bool, pos int) []string {
	var ones []string
	var zeroes []string

	var most []string
	var least []string

	for i := 0; i < len(input); i++ {
		chars := []rune(input[i])
		if string(chars[pos]) == "1" {
			ones = append(ones, input[i])
		} else {
			zeroes = append(zeroes, input[i])
		}
	}

	if len(ones) == len(zeroes) {
		if useMost {
			return ones
		} else {
			return zeroes
		}
	} else if len(ones) > len(zeroes) {
		most = ones
		least = zeroes
	} else {
		most = zeroes
		least = ones
	}

	if useMost {
		return most
	} else {
		return least
	}
}

func bitStringToInt(bits string) int {
	chars := []rune(bits)
	value := 0
	for i := 0; i < len(chars); i++ {
		pos := 11 - i
		if string(chars[i]) == "1" {
			value = setBit(value, pos)
		}
	}
	return value
}

func setBit(n int, pos int) int {
	n |= (1 << pos)
	return n
}
