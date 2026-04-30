package main

import (
	"fmt"
)

func countvowel(s string) int {
	count := 0
	// Option A: lowercase the whole string once
	//for _, r := range strings.ToLower(s) {
	for _, i := range s {
		switch i {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countvowel("hello"))
}
