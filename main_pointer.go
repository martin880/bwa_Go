package main

import "fmt"

func main() {
	number := 5
	fmt.Println("Alamat memory :", &number)
	fmt.Println("Nilai awal :", number)

	change(&number, 100)

	fmt.Println("Nilai Akhir :", number)
	fmt.Println("Alamat Memory :", &number)
}

func change(old *int, new int) {
	*old = new
	fmt.Println("Alamat memory :", old)
	fmt.Println("Di dalam function :", *old)
}
