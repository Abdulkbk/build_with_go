package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	bytes := flag.Bool("c", false, "Count Bytes")
	lines := flag.Bool("l", false, "Count Lines")
	words := flag.Bool("w", false, "Count Words")
	characters := flag.Bool("m", false, "Count Characters")
	filename := flag.String("file", "", "File name")

	// Parse the input flags provided by the user
	flag.Parse()

	// fmt.Println(*filename)
	// fmt.Println(*bytes)
	// fmt.Println(*lines)

	// get filename from cli
	// filenames := flag.Args()
	// fmt.Println(filenames)

	if *filename != "" {

		// fmt.Printf("Processing file: %s\n", *filename)
		readFile(*filename, *bytes, *lines, *words, *characters)
	}

}

func readFile(file string, bytes bool, lines bool, words bool, chars bool) {
	text, err := getContent(file)
	if !bytes && !lines && !words && !chars {
		c := count(text, true, lines, words, chars)
		l := count(text, bytes, true, words, chars)
		w := count(text, bytes, lines, true, chars)
		m := count(text, bytes, lines, words, true)

		fmt.Printf("%v  %v  %v  %v  %s\n", c, l, w, m, file)
		return
	}

	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", file, err)
	}

	value := count(text, bytes, lines, words, chars)

	fmt.Printf("%v %s\n", value, file)
}

func getContent(filename string) (io.Reader, error) {
	content, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(content), nil
}

func count(r io.Reader, countBytes bool, countLines bool, countWords bool, countChars bool) int {
	scanner := bufio.NewScanner(r)

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else if countLines {
		scanner.Split(bufio.ScanLines)
	} else if countWords {
		scanner.Split(bufio.ScanWords)
	} else if countChars {
		scanner.Split(bufio.ScanRunes)
	}
	result := 0

	for scanner.Scan() {
		result++
	}
	return result
}
