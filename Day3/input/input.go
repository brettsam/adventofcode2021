package input

import (
	"bufio"
	"log"
	"os"
)

func GetInput() []string {
	readFile, err := os.Open("input/input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var input []string

	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}

	readFile.Close()

	return input
}
