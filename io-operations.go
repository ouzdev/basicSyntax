package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	written, err := CopyFile("destination.txt", "source.txt")

	if err != nil {
		fmt.Println(err.Error())
	}
	//The "written" mean copied character count.
	fmt.Println(written)
}

// Let's look at a function that opens two files and copies the contents of one file to the other
func CopyFile(dstName, srcName string) (written int64, err error) {

	//Create file if not exists
	src, err := os.OpenFile(srcName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return 0, err
	}
	//Close file with defer keyword if occur error
	defer src.Close()

	//Create source file content
	var sourceContent = "Source Content 2"

	//Write content to source file
	_, err = io.WriteString(src, sourceContent)

	//Because WriteString is used, the cursor will be from the end of the content.
	//To do this, we move the cursor to the start with the seek method.
	src.Seek(0, 0)

	if err != nil {
		return 0, err
	}

	dst, err := os.OpenFile(dstName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return 0, err
	}

	defer dst.Close()

	//This here, we copy source content to destination file content
	written, err = io.Copy(dst, src)

	//Error handling
	if err != nil {
		return 0, err
	}

	//Return written
	return written, nil
}
