package main

import (
	"fmt"
	"time"
)

func consolePrint(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from)
	}
}

func main() {

	consolePrint("Simpra PMS")

	go consolePrint("Simpra Resort Manager 1")

	go consolePrint("Simpra Resort Manager 2")
	fmt.Println("İşleri bozan ben !!!!")
	go consolePrint("Simpra Resort Manager 3")
	fmt.Println("İşleri bozan ben !!!!")
	go consolePrint("Simpra Resort Manager 4")

	time.Sleep(time.Second)

	fmt.Println("Done!!!")

	/*
		Output:
			Simpra PMS
			Simpra PMS
			Simpra PMS
			İşleri bozan ben !!!!
			Simpra Resort Manager
			Simpra Resort Manager
			Simpra Resort Manager
			Done!!!
	*/
}
