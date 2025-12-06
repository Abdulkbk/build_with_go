package main

import "bufio"

// getCount is a generalize function to get the actual number of
// or counts of the element we are interested in.
func getCount(scanner *bufio.Scanner, scanType bufio.SplitFunc) int {
	count := 0
	scanner.Split(scanType)

	for scanner.Scan() {
		count++
	}

	return count
}
