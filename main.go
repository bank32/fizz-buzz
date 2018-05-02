package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i < 100; i++ {
		var msg string

		if i%3 == 0 {
			msg = "Fizz"
		} else if i%5 == 0 {
			msg = "Buzz"
		}

		fmt.Printf(strconv.Itoa(i)+" %v\n", msg)
	}
}
