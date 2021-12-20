package input

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func GetInput() map[string]string {
	readFile, err := os.Open("input/input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	nodeStrings := make(map[string]string, 0)
	for fileScanner.Scan() {
		content := fileScanner.Text()

		var inputLine = strings.Split(content, " -> ")
		nodeStrings[inputLine[0]] = inputLine[1]
	}

	readFile.Close()

	return nodeStrings
}
