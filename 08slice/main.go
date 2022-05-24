package main

import (
	"fmt"
	"sort"
)

func main() {

	chars := []string{"a", "b", "c"}
	fmt.Printf("chars type %T \n", chars)
	// chars[4] = "d" // error
	chars = append(chars, "d")
	fmt.Println(chars)
	chars = append(chars[1:2])
	fmt.Println(chars)

	nums := make([]int, 3)
	nums[0] = 1
	nums[1] = 78
	nums[2] = 16
	// nums[3] = 100 // error
	nums = append(nums, 600, 200)
	fmt.Println(nums)
	sort.Ints(nums)
	fmt.Println(nums)

	// delete by index
	index := 2
	nums = append(nums[:index], nums[index+1:]...)
	fmt.Println(nums)

}
