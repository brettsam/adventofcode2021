package input

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetMeasurements() []int {
	readFile, err := os.Open("input/input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var measurements []int

	for fileScanner.Scan() {
		m, _ := strconv.ParseInt(fileScanner.Text(), 10, 32)
		measurements = append(measurements, int(m))
	}

	readFile.Close()

	return measurements
}
