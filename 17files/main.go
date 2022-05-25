package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "this will be text content of a file we create"
	fileName := "./file.txt"
	file, err := os.Create(fileName)
	checkNillErr(err)

	length, err := io.WriteString(file, content)
	checkNillErr(err)

	fmt.Println(length)
	defer file.Close()
	readFile(fileName)
}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)

	checkNillErr(err)

	fmt.Println(string(databyte))
}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
