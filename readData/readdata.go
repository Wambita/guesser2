package readData

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"guess-it-2/calculations"
)

func Reader() {
	var numslice []float64
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		numInt := scanner.Text()
		numFloat, err := strconv.ParseFloat(numInt, 64)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		numslice = append(numslice, numFloat)
		if len(numslice) > 1 {
			lowRange, highRange := calculations.Range(numslice)
			fmt.Printf("%d %d\n", lowRange, highRange)
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
