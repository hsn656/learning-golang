package main

import "fmt"

func main() {
	var nums [3]int
	nums[0] = 1
	nums[1] = 2

	fmt.Println(nums)
	fmt.Println(len(nums))

	chars := [3]string{"a", "b", "c"}
	// chars[4] = "d" // compilation error
	fmt.Println(chars)
}
