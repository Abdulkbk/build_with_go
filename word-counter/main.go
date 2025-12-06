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

	sc := bufio.NewScanner(file)

	if *c {
		bytes := getCount(sc, bufio.ScanBytes)
		fmt.Printf("%v %s", bytes, filename)
		return
	}

	if *l {
		lines := getCount(sc, bufio.ScanLines)

		fmt.Printf("%v %s", lines, filename)
		return
	}

	if *w {
		words := getCount(sc, bufio.ScanWords)

		fmt.Printf("%v %s", words, filename)
		return
	}

	if *m {
		runes := getCount(sc, bufio.ScanRunes)

		fmt.Printf("%v %s", runes, filename)
		return
	}

}
