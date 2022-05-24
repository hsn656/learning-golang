package main

import "fmt"

func main() {
	loginCount := 26
	var result string
	if loginCount < 10 {
		result = "no pressure yet"
	} else if loginCount > 10 {
		result = "pressure"
	} else {
		result = "login count is 10, prepare"
	}
	fmt.Println(result)

	if num := 3; num&2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

}
