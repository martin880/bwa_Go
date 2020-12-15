package main

import "fmt"

func main() {
	students := []map[string]string{
		{"name": "Martin", "score": "A"},
		{"name": "Rossi", "score": "B"},
		{"name": "Marquez", "score": "E"},
	}

	for _, student := range students {
		fmt.Println(student["name"], " - ", student["score"])
	}
}
