package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	data, err := os.ReadFile("people.json")
	if err != nil {
		log.Fatal(err)
	}
	var people []Person
	if err := json.Unmarshal(data, &people); err != nil {
		log.Fatal(err)
	}

	for _, p := range people {
		fmt.Printf("Name: %s Age:%d \n", p.Name, p.Age)
	}
}
