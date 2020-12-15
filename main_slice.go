package main

import "fmt"

func main() {
	var gamingConsoles []string

	gamingConsoles = append(gamingConsoles, "Playstation 4")
	gamingConsoles = append(gamingConsoles, "Nintendo Switch")
	gamingConsoles = append(gamingConsoles, "XBox One")

	for _, console := range gamingConsoles {
		fmt.Println(console)
	}
}
