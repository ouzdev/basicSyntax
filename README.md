# Go Basic Syntax
![Go Workflow](https://github.com/ouzdev/basicSyntax/workflows/Go/badge.svg)


This is a simple Go program that prints "Hello, World!" to the console.

## Getting Started

To run this program on your local machine, follow these steps:

1. Make sure you have Go installed. If not, you can download it [here](https://golang.org/dl/).

2. Clone this repository:

    ```bash
    git clone https://github.com/ouzdev/basicSyntax.git
    ```

3. Navigate to the project directory:

    ```bash
    cd basicSyntax
    ```

4. Run the program:

    ```bash
    go run syntax.go
    ```

## Code Structure

The code is organized into the following sections:

- **Package Declaration:** Every Go program is part of a package. In this example, the program belongs to the main package.

- **Import Statement:** The `import "fmt"` statement imports the `fmt` package, which provides functions for formatting and printing.

- **Main Function:** The `main` function is the entry point of the program. Code inside its curly braces will be executed. Here, it prints "Hello, World!" to the console.

## Code Compact Example

There's also a compact version provided in the comments, but it's not recommended for readability.

```go
package main; import "fmt"; func main() {fmt.Println("Hello, World!")}
