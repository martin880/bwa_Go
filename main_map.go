package main

import "fmt"

func main() {
	myMap := map[string]string{
		"ruby":       "is beatifull",
		"go":         "is powerfull and superfast",
		"Javascript": "is hype",
	}
	for key, value := range myMap {
		fmt.Println("Key: ", key, " value: ", value)
	}

	fmt.Println("===================================")

	delete(myMap, "ruby")

	for key, value := range myMap {
		fmt.Println("Key: ", key, " value: ", value)
	}

	// mengecek sebuah value ada atau tidak
	// value, isAvailable := myMap["python"]
	// fmt.Println(isAvailable)
	// fmt.Println(value)
}
