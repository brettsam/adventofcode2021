package input

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInput() []int {
	readFile, err := os.Open("input/input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var numbers []int

	for fileScanner.Scan() {
		content := fileScanner.Text()
		var s = strings.Split(content, ",")
		for i := 0; i < len(s); i++ {
			num, _ := strconv.ParseInt(s[i], 10, 32)
			numbers = append(numbers, int(num))
		}
	}

	readFile.Close()

	return numbers
}
