package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File System In go lang")

	content := "Welcome to File Handling in Go Lagng ad"

	filePtr, err := os.Create("./myfile1.txt")
	checkNilError(err)

	// write Something in File
	len, err := io.WriteString(filePtr, content)
	checkNilError(err)

	fmt.Println("Lentgh of file: ", len)

	// close the file with defer
	defer filePtr.Close()

	// for Read the file
	readFile("./myfile1.txt")
}

func readFile(fileName string) {
	dataInBytes, err := os.ReadFile(fileName)
	checkNilError(err)

	fmt.Println("Raw Data Inside file: ", dataInBytes)
	fmt.Println("Actual Data Inside file: ", string(dataInBytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
