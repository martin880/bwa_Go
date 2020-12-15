package main

import "fmt"

func main() {

	// var languages [5]string
	// languages[0] = "Go"
	// languages[1] = "Ruby"
	// languages[2] = "JavaScript"
	// languages[3] = "Ruby"
	// languages[4] = "Phyton"

	languages := [...]string{
		"Go",
		"Ruby",
		"Javasript",
		"Ruby",
		"Phyton",
		"Kotlin",
	}

	//for dipakai untuk mengambil setiap elemen dalam array dengan range
	for index, lang := range languages {
		fmt.Println("Index: ", index, " languages : ", lang)
	}
}
