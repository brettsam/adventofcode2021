package input

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/brettsam/adventofcode2021/day11/octo"
)

func GetInput() [][]octo.Octo {
	readFile, err := os.Open("input/input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var numbers [][]octo.Octo

	for fileScanner.Scan() {
		content := fileScanner.Text()
		var inputLine = strings.Split(content, "")

		var intLine []octo.Octo
		for _, s := range inputLine {
			i, _ := strconv.ParseInt(s, 10, 32)
			newOcto := octo.Octo{Value: int(i), Flashed: false}
			intLine = append(intLine, newOcto)
		}

		numbers = append(numbers, intLine)
	}

	readFile.Close()

	return numbers
}
