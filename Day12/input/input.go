package input

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/brettsam/adventofcode2021/day12/node"
)

func GetInput() *node.Node {
	readFile, err := os.Open("input/input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var nodeStrings [][]string

	for fileScanner.Scan() {
		content := fileScanner.Text()
		var inputLine = strings.Split(content, "-")
		nodeStrings = append(nodeStrings, inputLine)
	}

	readFile.Close()

	m := make(map[string]*node.Node)

	for _, a := range nodeStrings {
		for _, b := range a {
			_, exists := m[b]
			if !exists {
				newNode := node.Node{Name: b}
				m[b] = &newNode
			}
		}
	}

	linkChildren(m, nodeStrings)

	return m["start"]
}

func linkChildren(m map[string]*node.Node, nodeStrings [][]string) {
	for key, value := range m {
		for _, a := range nodeStrings {
			for j, b := range a {
				if b == key {
					var childNodeName string
					if j == 0 {
						childNodeName = a[1]
					} else {
						childNodeName = a[0]
					}
					childNode := m[childNodeName]
					value.Nodes = append(value.Nodes, childNode)
				}
			}
		}
	}
}
