package main

import (
	"fmt"
)

func main() {

	sum := 0 // var sum int = 0

	for i := 0; i < 10; i++ {
		fmt.Println(i, ". --> ", i)
		sum += i
	}

	fmt.Println("Sum:", sum)

	/*
		  Output:
			0 . -->  0
			1 . -->  1
			2 . -->  2
			3 . -->  3
			4 . -->  4
			5 . -->  5
			6 . -->  6
			7 . -->  7
			8 . -->  8
			9 . -->  9
			Sum: 45
	*/
}
