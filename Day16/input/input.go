package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInput() string {
	readFile, err := os.Open("input/input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	bits := ""
	for fileScanner.Scan() {
		content := fileScanner.Text()
		for _, c := range strings.Split(content, "") {
			hex, _ := strconv.ParseInt(c, 16, 8)
			bitString := fmt.Sprintf("%04b", hex)
			bits += bitString
		}
	}

	readFile.Close()

	return bits
}
