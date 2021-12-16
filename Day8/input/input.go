package input

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/brettsam/adventofcode2021/day8/line"
)

func GetInput() []line.Line {
	readFile, err := os.Open("input/input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var numbers []line.Line

	for fileScanner.Scan() {
		content := fileScanner.Text()
		var split = strings.Split(content, " | ")
		var input [10]string
		var inputSlice = strings.Fields(split[0])
		for i, s := range inputSlice {
			inputSlice[i] = sortString(s)
		}
		copy(input[:], inputSlice)

		var output [4]string
		var outputSlice = strings.Fields(split[1])
		for i, s := range outputSlice {
			outputSlice[i] = sortString(s)
		}
		copy(output[:], outputSlice)

		var l = line.Line{Input: input, Output: output}
		numbers = append(numbers, l)
	}

	readFile.Close()

	return numbers
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
