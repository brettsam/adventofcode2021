package input

import (
	"bufio"
	"log"
	"os"
)

func GetDirections() []string {
	readFile, err := os.Open("input/input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var directions []string

	for fileScanner.Scan() {
		directions = append(directions, fileScanner.Text())
	}

	readFile.Close()

	return directions
}
