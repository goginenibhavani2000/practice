package main

import (
	"fmt"
	"strings"
)

func wordFreq (s string) map[string]int{
	wordMap := make(map[string]int)
	// can use strings.Split(s, " ") also but .Fields is better 
	for _,word := range strings.Fields(strings.ToLower(s)) {
		wordMap[word]++;
	}
	return wordMap
}

func main(){
	w := wordFreq("Go basics practice , go")
	for word, count := range w{
		fmt.Println("word is ", word , " freq is ", count)
	}
}