package main

import (
	"fmt"
	"unicode"

	"github.com/brettsam/adventofcode2021/day12/input"
	"github.com/brettsam/adventofcode2021/day12/node"
)

var routes int

func main() {
	numbers := input.GetInput()
	fmt.Println(partOne(numbers))

	numbers = input.GetInput()
	fmt.Println(partTwo(numbers))
}

func partOne(numbers *node.Node) int {
	findRoutes(numbers)
	return routes
}

func partTwo(numbers *node.Node) int {
	routes = 0
	r := make([]string, 0)
	findRoute2(numbers, r)
	return routes
}

func findRoutes(n *node.Node) {
	r := make([]string, 0)
	findRoute(n, r)
}

func findRoute(n *node.Node, curRoute stack) {
	curRoute = curRoute.Push(n.Name)
	for _, child := range n.Nodes {
		if child.Name == "start" {
			continue
		}
		if child.Name == "end" {
			curRoute = curRoute.Push("end")
			printRoute(curRoute)
			routes++
			curRoute, _ = curRoute.Pop()
			continue
		}
		if isLower(child.Name) && contains(curRoute, child.Name) {
			continue
		}
		findRoute(child, curRoute)
	}
}

func findRoute2(n *node.Node, curRoute stack) {
	curRoute = curRoute.Push(n.Name)
	for _, child := range n.Nodes {
		if child.Name == "start" {
			continue
		}
		if child.Name == "end" {
			curRoute = curRoute.Push("end")
			//printRoute(curRoute)
			routes++
			curRoute, _ = curRoute.Pop()
			continue
		}
		if isLower(child.Name) && contains(curRoute, child.Name) && anyDupes(curRoute) {
			continue
		}
		findRoute2(child, curRoute)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func anyDupes(s []string) bool {
	m := make(map[string]int, 0)
	for _, a := range s {
		if isLower(a) {
			m[a] = m[a] + 1
			if m[a] == 2 {
				return true
			}
		}
	}
	return false
}

func printRoute(r []string) {
	for _, s := range r {
		fmt.Print(s)
		if s != "end" {
			fmt.Print("->")
		} else {
			fmt.Println()
		}
	}
}

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}
