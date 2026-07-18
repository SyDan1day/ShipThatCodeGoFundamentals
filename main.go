package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	// Print FizzBuzz / Fizz / Buzz / n.
	switch {
	case n%5 == 0 && n%3 == 0:
		fmt.Println("FizzBuzz")
	case n%5 == 0 && n%3 != 0:
		fmt.Println("Buzz")
	case n%5 != 0 && n%3 == 0:
		fmt.Println("Fizz")
	default:
		fmt.Println(n)
	}
}
