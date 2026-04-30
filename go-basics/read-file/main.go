package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// read file line-by-line
// Key idea: bufio.Scanner for line-by-line. Always defer file.Close(). Always check scanner.Err() after the loop.

func main() {
	file, err := os.Open("path")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linenum := 1
	for scanner.Scan() {
		fmt.Printf("%d: %s \n", linenum, scanner.Text())
		linenum++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
