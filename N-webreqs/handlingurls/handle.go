package handlingurls

import (
	"fmt"
	"net/url"
)

const URL = "https://www.normalwebsite.com/learn?course=AIML&name=navnath"

func HandleSomeApi() {
	fmt.Println("Welcome to Handling Urls in Go Lang")

	result, err := url.Parse(URL)
	if err != nil {
		panic(err)
	}

	// fmt.Println("URLS Protocol: ", result.Scheme)
	// fmt.Println("URLS Path: ",result.Path)
	// fmt.Println("URL Queries: ",result.Query())
	// fmt.Println("URL Host: ",result.Host)
	// fmt.Println("URL Port: ",result.Port())

	fmt.Println("URL Queries: ", result.RawQuery)

	// its map[string] = []string
	params := result.Query()
	fmt.Printf("Type of q: %T", params) // Type url.Values it means map

	course := params["course"]
	fmt.Println("Course: ", course[0])
	fmt.Printf("Type of course: %T \n", course) // Type url.Values it means map

	partOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "www.ndkdev.me",
		Path:    "/projects",
		RawPath: "user=ndk",
	}

	anotherUrl := partOfUrl.String()
	fmt.Println("Custom URL: ", anotherUrl)
}
