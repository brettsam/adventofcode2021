package input

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetInput() [][]int {
	readFile, err := os.Open("input/input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var numbers [][]int

	for fileScanner.Scan() {
		content := fileScanner.Text()
		var inputLine = strings.Split(content, "")

		var intLine []int
		for _, s := range inputLine {
			i, _ := strconv.ParseInt(s, 10, 32)
			intLine = append(intLine, int(i))
		}

		numbers = append(numbers, intLine)
	}

	readFile.Close()

	return numbers
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
