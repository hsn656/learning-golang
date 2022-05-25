package main

import (
	"fmt"
	"net/url"
)

//some url from youtube
const myurl = "https://www.youtube.com:80/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	result, _ := url.Parse(myurl)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)

	qparams := result.Query()
	fmt.Printf("%T\n", qparams)
	fmt.Println(qparams["index"][0])

	for _, v := range qparams {
		fmt.Println(v)
	}
}
