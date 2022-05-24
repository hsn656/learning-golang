package main

import "fmt"

func main() {
	greet("hsn")

	func() {
		fmt.Println("IIEF")
	}()

	fmt.Println(sum(2, 3))
	result, _ := sumMany(2, 3, 4, 5)
	fmt.Println(result)

}

func greet(name string) {
	fmt.Println("hi", name)
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func sumMany(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}

	return total, "done"
}
