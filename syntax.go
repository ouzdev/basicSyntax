/*
	In Go, every program is part of a package. We define this using the package keyword!!

In this example, the program belongs to the main package.
*/
package main

/* import "fmt" lets us import files included it the fmt package.

If you want more learn more, visit this link https://pkg.go.dev/fmt

*/
import "fmt"

// If this file is not have main package when get error of main error

/*
	This is a function. Any code inside its curl brackets {} will be executed.

Note: In Go, any executable code belongs to the main package.
*/
func main() {
	//Println is a function made available from fmt package. It is used output/print text.
	fmt.Println("Hello World !")
}

/* Code Compact Example (Single line)
This is not recommended because it mnakes the code more diffucult to read
package main; import "fmt" func main() {fmt.Println("HelloWorld!")} */
