package main

import "fmt"

// net, entityFramework, signalR are boolean variables with default value false
// Structure --> var {bla bla bla bla} 'type'(bool etc.)
var net, entityFramework, signalR bool

func main() {
	// i is an integer variable with default value 0
	var i int

	// Print the values of i, net, signalR, entityFramework
	fmt.Println("Default values:", i, net, signalR, entityFramework)

	//Output: Default Values: 0 false false false
}
