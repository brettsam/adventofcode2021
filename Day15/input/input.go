package input

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInput() [][]*Node {
	readFile, err := os.Open("input/input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	nodes := make([][]*Node, 0)
	for fileScanner.Scan() {
		content := fileScanner.Text()

		lineString := strings.Split(content, "")
		var lineNodes []*Node
		for x, c := range lineString {
			v, _ := strconv.ParseInt(c, 10, 32)
			lineNodes = append(lineNodes, &Node{Risk: int(v), X: x, Y: len(nodes)})
		}
		nodes = append(nodes, lineNodes)
	}

	readFile.Close()

	return nodes
}

type Node struct {
	X    int
	Y    int
	Risk int
	Sum  int
}
