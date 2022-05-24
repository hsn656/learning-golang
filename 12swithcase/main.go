package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	num := rand.Intn(4) + 1

	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		fallthrough // to exec next case, no break here
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("not between 1-4")
	}

}
