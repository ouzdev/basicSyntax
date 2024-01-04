package main

import (
	"fmt"
	"log"
	"os"
)

// CustomError --> Struct is value type. Struct is not reference type!!!
// Go have not class type
type CustomError struct {
	Code    int
	Message string
}

type Customer struct {
	Name  string
	Price float32
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func main() {

	fmt.Println(Customer{Price: 150.58, Name: "Oğuz Can Genç"})

	f, err := os.OpenFile("filename.txt", os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Println(CustomError{Code: 404, Message: err.Error()})
	}

	str := f.Name()

	fmt.Println(str)
}
