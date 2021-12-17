package input

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func GetInput() [][]string {
	readFile, err := os.Open("input/input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var numbers [][]string

	for fileScanner.Scan() {
		content := fileScanner.Text()
		var inputLine = strings.Split(content, "")
		numbers = append(numbers, inputLine)
	}

	readFile.Close()

	return numbers
}
