package main

import (
	"bufio"
	"log"
	"os"
)

// getCount is a generalize function to get the actual number of
// or counts of the element we are interested in.
func getCount(filename string, scanType bufio.SplitFunc) int {
	count := 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}

	sc := bufio.NewScanner(file)
	sc.Split(scanType)

	for sc.Scan() {
		count++
	}

	return count
}
