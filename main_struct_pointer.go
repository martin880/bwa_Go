package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float64
}

// func main() {
// 	student := Student{1, "Martin Lumbangaol", 3.01}

// 	fmt.Println(student.Name)

// 	graduate(&student)

// 	fmt.Println(student.Name)
// }

// func graduate(student *Student) {
// 	student.Name = student.Name + " S.Kom"
// }

// method dengan pointer receiver
func (student *Student) graduate() {
	student.Name = student.Name + " S.Kom"
}

func main() {
	student := Student{1, "Martin Lumbangaol", 3.01}

	fmt.Println(student.Name)

	student.graduate()

	fmt.Println(student.Name)
}
