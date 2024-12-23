package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("datafile.txt")
	// defer file.Close()
	if err != nil {
		panic(err)
	}
	fileLength, err := io.WriteString(file, "String is good")
	if err != nil {
		panic(err)
	}
	fmt.Println("File length is: ", fileLength)
}
