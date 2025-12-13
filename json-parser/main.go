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

	var stack []string

	for scanner.Scan() {
		el := scanner.Text()

		if el == "\n" || el == " " {
			continue
		}

		if firstChar == "" {
			firstChar = el
		}

		lastChar = el

		if el == "{" || el == "," {
			stack = append(stack, el)

			continue
		}

		if el == ":" {
			lastElIdx := len(stack) - 1
			lastEl := stack[lastElIdx]

			if lastEl != "{" {
				fmt.Println(stack)
				log.Fatal("parsing failed: `:`")
			}

			stack = append(stack, el)

			continue
		}

		if el == "\"" {
			lastElIdx := len(stack) - 1
			lastEl := stack[lastElIdx]

			if lastEl == ":" || lastEl == "\"" || lastEl == "," {
				stack = append(stack[:lastElIdx], stack[lastElIdx+1:]...)
				continue
			}

			stack = append(stack, el)

			continue
		}

		if el == "}" && len(stack) > 1 {
			lastElIdx := len(stack) - 1
			stack = append(stack[:lastElIdx], stack[lastElIdx+1:]...)

			lastElIdx = len(stack) - 1
			lastEl := stack[lastElIdx]
			if lastEl != "{" {
				fmt.Println(lastEl)
				log.Fatal("invalid json")
			}

			stack = append(stack[:lastElIdx], stack[lastElIdx+1:]...)

			continue
		}
	}

	if firstChar != "{" || lastChar != "}" {
		log.Fatalf("invalid first char(%v) and last char(%v)", firstChar, lastChar)
	}

	fmt.Println("valid json")

}
