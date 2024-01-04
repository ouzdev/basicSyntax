package main

import (
	"fmt"
	"math"
)

func main() {

	var a bool = true              // Boolean
	var b int = 5                  //Integer
	var c float32 = 3.854785659685 //Floating point number
	var d string = "Hi!"           //String
	var e float64 = math.MaxFloat64
	var f float64 = 6.25874587458

	//Declare multiple variables
	var (
		ToBe   bool   = false
		MaxInt uint64 = 1<<64 - 1
	)

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float 32 ", c)
	fmt.Println("String:  ", d)
	fmt.Println("Max Float 64:  ", e)
	fmt.Println("Float 64 ", f)

	fmt.Println(ToBe, MaxInt)
	/* Output

		Boolean:  true
	    Integer:  5
	    Float:    3.14
	    String:   Hi!
	*/
}
