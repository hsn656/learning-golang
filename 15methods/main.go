package main

import "fmt"

func main() {
	amr := User{Name: "amr", Email: "amr@hsn.com", age: 23, IsVerified: true}

	fmt.Println(amr.getAge())

	amr.setAge(26)

	fmt.Println(amr.getAge()) // no change as it uses copy only

}

type User struct {
	Name       string
	Email      string
	age        int
	IsVerified bool
}

func (u User) getAge() int {
	return u.age
}

func (u User) setAge(value int) {
	u.age = value
}
