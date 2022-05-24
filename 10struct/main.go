package main

import "fmt"

func main() {
	hsn := User{"hsn", "hsn@hsn.com", 23, true}
	amr := User{Name: "amr", Email: "amr@hsn.com", Age: 23, IsVerified: true}

	fmt.Println(hsn)
	fmt.Println(amr)
	fmt.Printf("%+v \n", amr)

	fmt.Println(hsn.Email)
	fmt.Printf("%v \n", amr.Age)

}

type User struct {
	Name       string
	Email      string
	Age        int
	IsVerified bool
}
