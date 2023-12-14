package main

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {

	t.Run("it should open existing file", func(t *testing.T) {
		_, err := getContent("test.txt")

		if err != nil {
			t.Fatalf("Test failed, `test.txt` file opened")
		} else {
			fmt.Println("Test passed: file `test.txt` opened")
		}
	})

	t.Run("It should fail if file does not exist", func(t *testing.T) {
		value, err := getContent("test1.txt")

		if err == nil {
			t.Fatalf("It should failed since `test1.txt` does not exist")

		} else {
			fmt.Printf("Test Passed: File `test1.txt` does not exist, value: %v", value)
		}
	})
}
