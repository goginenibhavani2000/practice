package main

import (
	"fmt"
)

func merge(s1 []int, s2 []int) []int {
	result := make([]int, 0, len(s1)+len(s2))
	i := 0
	j := 0
	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			result = append(result, s1[i])
			i++
		} else {
			result = append(result, s2[j])
			j++
		}
	}
	// if(i!=len(s1)){
	// 	result = append(result, s1[i])
	// 	i++
	// }
	// if(j!=len(s2)){
	// 	result = append(result, s2[j])
	// 	j++
	// }

	result = append(result, s1[i:]...) // append rest of a (if any)
	result = append(result, s2[j:]...) // append rest of b (if any)
	return result
}
func main() {
	fmt.Println(merge([]int{1, 3, 5}, []int{2, 4, 6})) // [1 2 3 4 5 6]
}
