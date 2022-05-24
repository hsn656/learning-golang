package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	fmt.Println("=================")
	for k, v := range nums {
		fmt.Println(k, v)
	}

	// works like while
	sum := 0
	for sum < 100 {
		sum += 30
		fmt.Println("sum now:", sum)
		if sum == 91 {
			goto somelabel
		}
	}

somelabel:
	fmt.Println("90 is enugh")
}
