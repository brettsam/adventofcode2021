package main

import (
	"fmt"
	"sort"

	"github.com/brettsam/adventofcode2021/day10/input"
)

func main() {
	numbers := input.GetInput()

	fmt.Println(partOne(numbers))
	fmt.Println(partTwo(numbers))
}

func partOne(numbers [][]string) int {
	count := 0

	for a := range numbers {
		s := make(stack, 0)
		err := ""
		match := true
		for b := range numbers[a] {
			cur := numbers[a][b]
			if cur == "(" || cur == "{" || cur == "<" || cur == "[" {
				s = s.Push(cur)
				continue
			} else if cur == ")" {
				s, match = validate(s, cur, "(")
			} else if cur == "}" {
				s, match = validate(s, cur, "{")
			} else if cur == ">" {
				s, match = validate(s, cur, "<")
			} else if cur == "]" {
				s, match = validate(s, cur, "[")
			}

			if !match {
				err = cur
				break
			}
		}
		count += addError(err)
	}

	return count
}

func partTwo(numbers [][]string) int {
	var scores []int
	for a := range numbers {
		s := make(stack, 0)
		err := ""
		match := true
		for b := range numbers[a] {
			cur := numbers[a][b]
			if cur == "(" || cur == "{" || cur == "<" || cur == "[" {
				s = s.Push(cur)
				continue
			} else if cur == ")" {
				s, match = validate(s, cur, "(")
			} else if cur == "}" {
				s, match = validate(s, cur, "{")
			} else if cur == ">" {
				s, match = validate(s, cur, "<")
			} else if cur == "]" {
				s, match = validate(s, cur, "[")
			}

			if !match {
				err = cur
				break
			}
		}
		// it's just incomplete!
		if err == "" {
			var seq []string
			rowCount := 0
			for i := len(s); i > 0; i-- {
				var c string
				s, c = s.Pop()
				seq = append(seq, getCloser(c))
			}

			for _, h := range seq {
				rowCount *= 5
				rowCount += getPoints(h)
			}
			scores = append(scores, rowCount)
		}
	}
	sort.Ints(scores)
	mid := (len(scores) - 1) / 2
	return scores[mid]
}

func getPoints(c string) int {
	switch c {
	case ")":
		return 1
	case "}":
		return 3
	case "]":
		return 2
	case ">":
		return 4
	default:
		return 0
	}
}
func getCloser(c string) string {
	switch c {
	case "(":
		return ")"
	case "{":
		return "}"
	case "[":
		return "]"
	case "<":
		return ">"
	default:
		return ""
	}
}

func addError(c string) int {
	switch c {
	case ")":
		return 3
	case "}":
		return 1197
	case ">":
		return 25137
	case "]":
		return 57
	default:
		return 0
	}
}

func validate(s stack, c string, expected string) (stack, bool) {
	s, m := s.Pop()
	match := true
	if m != expected {
		match = false
	}
	return s, match
}

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}
