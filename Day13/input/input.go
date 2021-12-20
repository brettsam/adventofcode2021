package input

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInput() [][]string {
	readFile, err := os.Open("input/input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var nodeStrings [][]string
	maxX := 0
	for fileScanner.Scan() {
		content := fileScanner.Text()
		if content == "" {
			break
		}
		var inputLine = strings.Split(content, ",")
		cx, _ := strconv.ParseInt(inputLine[0], 10, 32)
		ix := int(cx)
		if ix > maxX {
			maxX = ix
		}
		cy, _ := strconv.ParseInt(inputLine[1], 10, 32)
		iy := int(cy)
		for x := 0; x <= ix; x++ {
			for y := 0; y <= iy; y++ {
				if y > len(nodeStrings)-1 {
					newLine := make([]string, cx)
					nodeStrings = append(nodeStrings, newLine)
				}
				if x > len(nodeStrings[y])-1 {
					nodeStrings[y] = append(nodeStrings[y], ".")
				}
				if nodeStrings[y][x] != "#" {
					nodeStrings[y][x] = "."
				}
			}
		}
		nodeStrings[iy][ix] = "#"
	}

	readFile.Close()

	for y, a := range nodeStrings {
		l := len(a) - 1
		for i := l; i < maxX; i++ {
			nodeStrings[y] = append(nodeStrings[y], ".")
		}
	}

	return nodeStrings
}
