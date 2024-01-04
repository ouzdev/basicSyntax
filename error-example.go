package main

import (
	"fmt"
	"log"
	"os"
)

type CustomError struct {
	Code    int
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func main() {

	f, err := os.ReadFile("filename.txt")

	if err != nil {
		log.Println(CustomError{Code: 404, Message: err.Error()})
	}

	str := string(f)

	fmt.Println(str)
}
