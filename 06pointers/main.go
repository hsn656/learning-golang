package main

import "fmt"

func main() {
	var ptr1 *int
	fmt.Println(ptr1)

	num := 3
	ptr2 := &num
	fmt.Println(ptr2)
	fmt.Println(*ptr2)

	*ptr2 = *ptr2 + 1
	fmt.Println(ptr2)
	fmt.Println(num)

}
