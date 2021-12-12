package input

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/brettsam/adventofcode2021/day5/point"
)

func GetInput() []point.Vertex {
	readFile, err := os.Open("input/input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var numbers []point.Vertex

	for fileScanner.Scan() {
		content := fileScanner.Text()
		var split = strings.Split(content, " -> ")
		var points1 = strings.Split(split[0], ",")
		var points2 = strings.Split(split[1], ",")

		point1X, _ := strconv.ParseInt(points1[0], 10, 32)
		point1Y, _ := strconv.ParseInt(points1[1], 10, 32)
		point1 := point.Point{int(point1X), int(point1Y)}

		point2X, _ := strconv.ParseInt(points2[0], 10, 32)
		point2Y, _ := strconv.ParseInt(points2[1], 10, 32)
		point2 := point.Point{int(point2X), int(point2Y)}

		vertex := point.Vertex{point1, point2}
		numbers = append(numbers, vertex)
	}

	readFile.Close()

	return numbers
}
