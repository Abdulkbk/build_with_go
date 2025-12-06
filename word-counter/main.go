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

	flag.Parse()

	scanner := bufio.NewScanner(file)

	if *c {
		bytes := 0
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			bytes++
		}

		fmt.Printf("%v %s", bytes, filename)
		return
	}

	if *l {
		lines := 0
		for scanner.Scan() {
			lines++
		}

		fmt.Printf("%v %s", lines, filename)
		return
	}

	if *w {
		words := 0
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}

		fmt.Printf("%v %s", words, filename)
		return
	}

}
