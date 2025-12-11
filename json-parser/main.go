package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		log.Fatal("file not specified")
	}

	filename := args[0]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanBytes)

	var firstChar string
	var lastChar string

	for scanner.Scan() {
		if firstChar == "" {
			firstChar = scanner.Text()
		}
		lastChar = scanner.Text()
	}

	if firstChar != "{" || lastChar != "}" {
		log.Fatalf("invalid first char(%v) and last char(%v)", firstChar, lastChar)
	}

	fmt.Println("valid json")

}
