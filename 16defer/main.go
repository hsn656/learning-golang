package main

import "fmt"

func main() {
	// it works like LIFO
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")

	defer deferer()
	deferer()
	fmt.Println("this will run first")

}

func deferer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
