package main

import (
	"fmt"
	"strconv"
)

func tell(no int) string {
	if no%3 == 0 && no%5 == 0 {
		return "FizzBuzz"
	} else if no%3 == 0 {
		return "Fizz"
	} else if no%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(no)
}

func main() {
	for i := 1; i < 100; i++ {
		fmt.Println(tell(i))
	}
}
