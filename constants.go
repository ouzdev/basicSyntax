package main

import (
	"fmt"
	"math"
)

const roomRate string = "Const Room Rate"

func main() {

	fmt.Println(roomRate)

	const roomRate = 50000

	const calculatedRoomRate = 3e20 / roomRate

	fmt.Println(calculatedRoomRate)

	fmt.Println(int64(calculatedRoomRate))

	//math.Sin expects a float64
	fmt.Println(math.Sin(roomRate))

	/* Output

	Const Room Rate
	6e+15
	6000000000000000
	-0.9998401890897897

	*/
}
