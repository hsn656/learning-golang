package main

import (
	"fmt"
	"time"
)

func main() {
	// 01/02/2006 => standard date
	// 15:04:05   => standard time
	// Monday     => standard day
	// check pkg doc for more
	fmt.Println(time.Now().Format("01/02/2006 Monday 15:04:05"))

	time.Sleep(3 * time.Second)
go 
	birthDate := time.Date(1996, time.May, 8, 0, 0, 0, 0, time.UTC)
	fmt.Println(birthDate.Format("01/02/2006 Monday"))

}
