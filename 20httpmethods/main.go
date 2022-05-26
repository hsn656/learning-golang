// launch the node serve in util folder to execute this properly

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest()
	//PerformPostRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(databyte))

	// another common way to read res data
	var responseString strings.Builder
	responseString.Write(databyte)
	fmt.Println(responseString.String())

}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"name":"hsn"
		}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	databyte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(databyte))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("name", "hsn")

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	databyte, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(databyte))

}
