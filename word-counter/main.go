package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "test.txt"

	_, err := os.Open(filename)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	c := flag.Bool("c", false, "count bytes")
	// l := flag.Bool("l", false, "count lines")

	flag.Parse()

	if *c {
		file, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}

		fmt.Printf("%v %s", len(file), filename)
	}

	// if *l {
	// 	lines := 0
	// 	for i := range file {

	// 	}

	// 	fmt.Printf("%v %v", lines, filename)
	// }
}
