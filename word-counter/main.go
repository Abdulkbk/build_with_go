package main

import (
	"bufio"
	"flag"
	"fmt"
)

func main() {
	var filename string

	c := flag.Bool("c", false, "count bytes")
	l := flag.Bool("l", false, "count lines")
	w := flag.Bool("w", false, "count words")
	m := flag.Bool("m", false, "count runes")

	flag.Parse()

	if *c {
		bytes := getCount(bufio.ScanBytes)
		fmt.Printf("%v %s", bytes, filename)
		return
	}

	if *l {
		lines := getCount(bufio.ScanLines)

		fmt.Printf("%v %s", lines, filename)
		return
	}

	if *w {
		words := getCount(bufio.ScanWords)

		fmt.Printf("%v %s", words, filename)
		return
	}

	if *m {
		runes := getCount(bufio.ScanRunes)

		fmt.Printf("%v %s", runes, filename)
		return
	}

	bytes := getCount(bufio.ScanBytes)
	lines := getCount(bufio.ScanLines)
	words := getCount(bufio.ScanWords)
	runes := getCount(bufio.ScanRunes)

	fmt.Printf("%v %v %v %v %s", bytes, lines, words, runes, filename)
}
