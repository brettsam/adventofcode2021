package main

import (
	"fmt"
	"strconv"

	"github.com/brettsam/adventofcode2021/day16/input"
)

func main() {
	start := input.GetInput()
	p1 := partOne(start)

	fmt.Println("--------------------")
	fmt.Println()

	start = input.GetInput()
	p2 := partTwo(start)

	fmt.Println()
	fmt.Println(p1)
	fmt.Println(p2)
}

var versionSum int

func partOne(bits string) int {
	parseHeader(bits)
	return versionSum
}

func partTwo(bits string) int64 {
	_, i := parseHeader(bits)
	return i
}

func parseHeader(bits string) (string, int64) {
	isLiteral := false
	v := bits[:3]
	bits = bits[3:]
	version, _ := strconv.ParseInt(v, 2, 8)
	versionSum += int(version)
	fmt.Println("VERSION", v, version)

	t := bits[:3]
	bits = bits[3:]
	tp, _ := strconv.ParseInt(t, 2, 8)
	typeString := ""
	var op func([]int64) int64
	switch tp {
	case 0:
		typeString = "SUM"
		op = sum
	case 1:
		typeString = "MULT"
		op = mult
	case 2:
		typeString = "MIN"
		op = min
	case 3:
		typeString = "MAX"
		op = max
	case 4:
		typeString = "LITERAL"
		isLiteral = true
	case 5:
		typeString = "GT"
		op = gt
	case 6:
		typeString = "LT"
		op = lt
	case 7:
		typeString = "EQUAL"
		op = equal
	}

	fmt.Println(typeString, t, tp)

	if isLiteral {
		return parseLiteral(bits)
	} else {
		t = bits[:1]
		bits = bits[1:]
		tp, _ := strconv.ParseInt(t, 2, 8)

		var nums []int64
		if tp == 0 {
			bits, nums = parseOperatorLength(bits)
		} else {
			bits, nums = parseOperatorNumber(bits)
		}
		return bits, op(nums)
	}
}

func parseOperatorLength(bits string) (string, []int64) {
	fmt.Println("LENGTH", 0, 0)
	t := bits[:15]
	bits = bits[15:]
	length, _ := strconv.ParseInt(t, 2, 64)
	fmt.Println("VALUE", t, length)
	fmt.Println()
	nums := make([]int64, 0)
	processed := 0
	prevBits := len(bits)
	for {
		var result int64
		bits, result = parseHeader(bits)
		nums = append(nums, result)
		processed += prevBits - len(bits)
		prevBits = len(bits)
		if processed >= int(length) {
			break
		}
	}
	return bits, nums
}

func parseOperatorNumber(bits string) (string, []int64) {
	fmt.Println("NUMBER", 1, 1)
	t := bits[:11]
	bits = bits[11:]
	number, _ := strconv.ParseInt(t, 2, 64)
	fmt.Println("VALUE", t, number)
	fmt.Println()
	nums := make([]int64, 0)
	for i := 0; i < int(number); i++ {
		var result int64
		bits, result = parseHeader(bits)
		nums = append(nums, result)
	}
	return bits, nums
}

func parseLiteral(bits string) (string, int64) {
	value := ""
	for {
		data := bits[0:5]
		bits = bits[5:]
		value += data[1:5]
		if data[0:1] == "0" {
			break
		}
	}
	digit, _ := strconv.ParseInt(value, 2, 64)
	fmt.Println("VALUE", value, digit)
	fmt.Println()
	return bits, digit
}

func sum(operands []int64) int64 {
	r := int64(0)
	for _, i := range operands {
		r += i
	}
	return r
}

func mult(operands []int64) int64 {
	r := int64(1)
	for _, i := range operands {
		r *= i
	}
	return r
}

func min(operands []int64) int64 {
	r := operands[0]
	for _, i := range operands {
		if i < r {
			r = i
		}
	}
	return r
}

func max(operands []int64) int64 {
	r := operands[0]
	for _, i := range operands {
		if i > r {
			r = i
		}
	}
	return r
}

func gt(operands []int64) int64 {
	if operands[0] > operands[1] {
		return 1
	}
	return 0
}

func lt(operands []int64) int64 {
	if operands[0] < operands[1] {
		return 1
	}
	return 0
}

func equal(operands []int64) int64 {
	if operands[0] == operands[1] {
		return 1
	}
	return 0
}
