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

	flag.Parse()

	if *c {
		file, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}

		fmt.Printf("%v %s", len(file), filename)
		return
	}

	if *l {
		lines := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines++
		}

		fmt.Printf("%v %s", lines, filename)
		return
	}

}
