package main

import (
	"fmt"
)

//Perulangan for
func main() {
	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("Saya sedang belajar Golang: ", i)
	// }

	// i := 1
	// for i <= 100 {
	// 	fmt.Println("Saya sedang belajar Go:", i)
	// 	i++
	// }

	title := "Golang the best language"

	// Program Go menampilkan huruf berdasarkan index genap
	// for index, letter := range title {
	// 	if index%2 == 0 {
	// 		fmt.Println("index :", index, "letter :", string(letter))
	// 	}
	// }

	// Program Go menampilkan huruf berdasarkan huruf vokal
	for index, letter := range title {
		letterString := string(letter)

		// if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
		// 	fmt.Println("index :", index, "letter :", string(letter))
		// }

		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("index :", index, "letter :", string(letter))
		}
	}

}

// func main() {
// 	// age := 9

// 	// if age > 10 {
// 	// 	fmt.Println("Kamu boleh bermain game")
// 	// } else {
// 	// 	fmt.Println("Kamu belum boleh bermain game")
// 	// }

// 	// score := 65
// 	// var grade string

// 	// if score <= 50 {
// 	// 	grade = "E"
// 	// } else if score <= 60 {
// 	// 	grade = "D"
// 	// } else if score <= 70 {
// 	// 	grade = "C"
// 	// } else {
// 	// 	grade = "NULL"
// 	// }

// 	// fmt.Println(grade)

// 	// switch case conditions
// 	// number := 5

// 	// switch number {
// 	// case 1:
// 	// 	fmt.Println("Satu")
// 	// case 2:
// 	// 	fmt.Println("Dua")
// 	// case 3:
// 	// 	fmt.Println("Tiga")
// 	// default:
// 	// 	fmt.Println("Tidak dikenali")
// 	// }
// }
