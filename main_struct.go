package main

import (
	"bwa_Go/management"
	"fmt"
)

// func displayUser(user User) string {
// 	return fmt.Sprintf("Name : %s %s, Email : %s", user.Firstname, user.Lastname, user.Email)
// }

// func displayGroup(group Group) {
// 	fmt.Printf("Name: %s", group.Name)
// 	fmt.Println("")
// 	fmt.Printf("Member count : %d", len(group.Users))
// 	fmt.Println("")

// 	fmt.Println("Users name :")
// 	for _, user := range group.Users {
// 		fmt.Println(user.Firstname)
// 	}
// }

//menggunakan konsep Exported pada package management
func main() {
	user := management.User{1, "Martin", "Ju", "martinju@gmail.com", true}
	result := user.Display()
	fmt.Println(result)

	user2 := management.User{2, "Iska", "Jenong", "iskajenong@gmail.com", true}
	fmt.Println(user2.Display())

	users := []management.User{user, user2}

	group := management.Group{"Gamer", user, users, true}

	// displayGroup(group)
	group.DisplayGroup()
}
