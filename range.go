package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

var marketPlaceItems = []string{"Police", "Genderme", "Barboon", "Simpra Link", "Protal"}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	/*
		Output:
		2**0 = 1
		2**1 = 2
		2**2 = 4
		2**3 = 8
		2**4 = 16
		2**5 = 32
		2**6 = 64
		2**7 = 128
	*/

	for index, item := range marketPlaceItems {
		// %d --> use with integer
		// %s --> use with string
		fmt.Printf("Id: %d, Item: %s\n", index, item)
	}

	/*
		Output:
		Id: 0, Item: Police
		Id: 1, Item: Genderme
		Id: 2, Item: Barboon
		Id: 3, Item: Simpra Link
		Id: 4, Item: Protal
	*/

}
