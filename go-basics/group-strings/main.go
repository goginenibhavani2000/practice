package main

import "fmt"

// group by length
func groupByLength(s []string) map[int][]string {
	stringslen := make(map[int][]string)
	for _, str := range s {
		// Key idea: append on a nil slice returns a new slice — so groups[k] = append(groups[k], v) works even when k isn't in the map yet.
		stringslen[len(str)] = append(stringslen[len(str)], str)
	}
	return stringslen
}

func main() {
	result := groupByLength([]string{"go", "rust", "c", "java", "ai"})
	for len, words := range result {
		fmt.Printf("len : %d strings: %v \n", len, words)
	}
}
