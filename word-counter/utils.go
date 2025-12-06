package main

import (
	"bufio"
	"log"
	"os"
)

// getCount is a generalize function to get the actual number of
// or counts of the element we are interested in.
func getCount(scanType bufio.SplitFunc) int {
	var (
		file *os.File
		err  error
	)

	args := os.Args

	switch {
	case len(args) > 1:
		filename := args[1]
		file, err = os.Open(filename)
		if err != nil {
			log.Fatalf("open file err: %v", err)
		}
		defer file.Close()

	default:
		file = os.Stdin
	}

	count := 0

	sc := bufio.NewScanner(file)
	sc.Split(scanType)

	for sc.Scan() {
		count++
	}

	return count
}
