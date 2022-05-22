package main

import "fmt"

// password := "password" // only can be used inside main

const Host = "localhost" // first letter is capital so it's public

func main() {
	var username string = "hsn"
	var age int = 26
	var isLogged = false
	password := "password"
	// password := "new password"  // error

	var (
		x = 1
		y = 2
	)

	fmt.Println(username, age)
	fmt.Printf("isLogged type:%T its value: %t \n", isLogged, isLogged)
	println(password) // i dont know why it worked withoud fmt
	println(x, y)
}
