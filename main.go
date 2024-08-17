package main

import (
	"fmt"
	"os"

	"guess-it-2/readData"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("usage: go run . ")
		return
	}
	readData.Reader()
}
