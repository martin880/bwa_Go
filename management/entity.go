package management

import "fmt"

//User struct for struct user
type User struct {
	ID        int
	Firstname string
	Lastname  string
	Email     string
	IsActive  bool
}

func (user User) Display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.Firstname, user.Lastname, user.Email)
}

//Group struct for struct group
type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func (group Group) DisplayGroup() {
	fmt.Printf("Name: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users name :")
	for _, user := range group.Users {
		fmt.Println(user.Firstname)
	}
}
