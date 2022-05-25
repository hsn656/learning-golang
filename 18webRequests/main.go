package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev/"

func main() {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%T\n", response)

	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(databyte))

	defer response.Body.Close()
}
