package main

import (
	"encoding/json"
	"log"
	"os"
)

// Write a struct slice to JSON
// Key idea: Struct tags control JSON field names. json.NewEncoder streams to a writer (better than Marshal + WriteFile for large data).
// we can use json.Marshal also
type Person struct{
	Name string
	Age int
}

func main() {
    people := []Person{
        {Name: "Alice", Age: 30},
        {Name: "Bob", Age: 25},
    }

	file, err:= os.Create("people.json")
	if err!=nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	if err=encoder.Encode(people); err!=nil{
		log.Fatal(err)
	}


}