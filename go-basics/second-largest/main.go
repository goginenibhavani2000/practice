package main

import (
	"fmt"
	"math"
)

func main() {
	// find second-largest number in slice
	input := []int{3, 1, 4, 1, 5, 9, 2, 6}
	if len(input) < 2 {
		fmt.Println(0)
	}
	var largest int = math.MinInt
	var second int = math.MinInt
	for _, v := range input {
		if v > largest {
			second, largest = largest, v
		} else if v > second && v != largest {
			second = v
		}
	}
	fmt.Print(second)

}
