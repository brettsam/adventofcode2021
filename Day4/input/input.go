package input

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInput() ([]int, [][5][5]int) {
	readFile, err := os.Open("input/input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var numbers []int
	var boards [][5][5]int

	isFirstLine := true
	currentLine := 0

	for fileScanner.Scan() {
		if isFirstLine {
			numbers = loadArray(fileScanner.Text(), ",")
			isFirstLine = false
			continue
		}

		content := fileScanner.Text()

		if content == "" {
			currentLine = 0
			var newBoard [5][5]int
			boards = append(boards, newBoard)
			continue
		}

		var currentBoard = &(boards[len(boards)-1])
		arrContent := loadArray(content, " ")
		for a := 0; a < len(arrContent); a++ {
			currentBoard[currentLine][a] = arrContent[a]
		}

		currentLine++
	}

	readFile.Close()

	return numbers, boards
}

func loadArray(input string, delim string) []int {
	var values []int

	var stringValues []string
	if delim == " " {
		stringValues = strings.Fields(input)
	} else {
		stringValues = strings.Split(input, delim)
	}

	for i := 0; i < len(stringValues); i++ {
		v, _ := strconv.ParseInt(stringValues[i], 10, 32)
		values = append(values, int(v))
	}
	return values
}
