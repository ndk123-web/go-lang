package main

import (
	"fmt"
	"io"
	"net/http"
	"webreq/handlingurls"
)

const url = "https://www.normalwebsite.com/"

func main() {
	fmt.Println("WEB REQ IN GO LANG")
	response, err := http.Get(url)
	CheckNilError(err)

	// Type of Response is Res
	fmt.Printf("Type %T", response)

	// callers responsibility to close the connection
	defer response.Body.Close()

	// binary stream -> io.Reader() -> bytes[] -> string 
	dataBytes, err := io.ReadAll(response.Body)
	CheckNilError(err)

	fmt.Println("Raw Data Bytes ", dataBytes)
	fmt.Println("Actual Data Bytes ", string(dataBytes))

	handlingurls.HandleSomeApi()
}

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}
