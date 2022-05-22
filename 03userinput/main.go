package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter your name")
	name, _ := reader.ReadString('\n')
	println("welcome, ", name)
}
