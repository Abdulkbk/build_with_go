package main

import (
	"bufio"
	"flag"
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

	reader := os.Stdin
	stat, err := reader.Stat()
	if err != nil {
		log.Fatalf("error reading stats: %v", err)
	}

	args := flag.Args()

	switch {
	case len(args) > 0:
		filename := args[0]
		file, err = os.Open(filename)
		if err != nil {
			log.Fatalf("open file err: %v", err)
		}
		defer file.Close()

	case stat.Size() > 0:
		file = os.Stdin

	default:
		log.Fatal("no source provided")
	}

	count := 0

	sc := bufio.NewScanner(file)
	sc.Split(scanType)

	for sc.Scan() {
		count++
	}

	return count
}
