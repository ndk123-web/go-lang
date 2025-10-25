package main

import (
	"fmt"
)

func main (){
	fmt.Println("Maps in Go Lang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["py"] = "Python"
	languages["cpp"] = "C Plus Plus"

	fmt.Println("All: ",languages)

	fmt.Println("JS: ",languages["JS"]) // it presents 
	fmt.Println("opas: ",languages["opas"]) // it not presents gives nothing  

	// delete any key from map
	delete(languages,"JS")
	fmt.Println("All: ",languages)

	// loop in maps
	for key , value := range languages{
		fmt.Println("Key: ",key, "value: ",value)
	}
}