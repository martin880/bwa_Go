package main

import "fmt"

func main() {
	// sentence := printMyResult("Saya sedang")
	// fmt.Println(sentence)
	// result := add(10, 20)
	// fmt.Println(result)
	luas, keliling := calculate(10, 2)
	fmt.Println(luas)
	fmt.Println(keliling)
}

func calculate(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}

// func add(number, numberTwo int) int {
// 	return number + numberTwo
// }

// func printMyResult(sentence string) string {
// 	newSentence := sentence + " belajar Golang"
// 	return newSentence
// }
