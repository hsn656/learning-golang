package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter number to increment:")
	input, _ := reader.ReadString('\n')
	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("not valid number")
	} else {
		fmt.Println(num + 1)
	}
}
