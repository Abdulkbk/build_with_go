package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "test.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	defer file.Close()

	c := flag.Bool("c", false, "count bytes")
	l := flag.Bool("l", false, "count lines")
	w := flag.Bool("w", false, "count words")
	m := flag.Bool("m", false, "count runes")

	flag.Parse()

	if *c {
		bytes := getCount(filename, bufio.ScanBytes)
		fmt.Printf("%v %s", bytes, filename)
		return
	}

	if *l {
		lines := getCount(filename, bufio.ScanLines)

		fmt.Printf("%v %s", lines, filename)
		return
	}

	if *w {
		words := getCount(filename, bufio.ScanWords)

		fmt.Printf("%v %s", words, filename)
		return
	}

	if *m {
		runes := getCount(filename, bufio.ScanRunes)

		fmt.Printf("%v %s", runes, filename)
		return
	}

	bytes := getCount(filename, bufio.ScanBytes)
	lines := getCount(filename, bufio.ScanLines)
	words := getCount(filename, bufio.ScanWords)
	runes := getCount(filename, bufio.ScanRunes)

	fmt.Printf("%v %v %v %v %s", bytes, lines, words, runes, filename)
}
