package waitgroups

import (
	"fmt"
	"net/http"
	"sync"
)

// create WaitGroup
// Its Advance version of Time
// because to wait the main routine we need some time to wait until other goroutine completed
var WaitGroupsList sync.WaitGroup // pointers

func WaitGroupsExample() {
	websiteList := []string{
		"https://google.com",
		"https://github.com",
		"https://x.com",
	}

	for _, v := range websiteList {

		// We created lightweight thread (goroutine)
		WaitGroupsList.Add(1)
		go GetStatusCode(v)
	}

	// This Function will wait till all routines sends Done as in below GetStatusCode
	WaitGroupsList.Wait()

}

func GetStatusCode(endpoint string) {

	// Responsibility to return DONE when Routine has completed
	defer WaitGroupsList.Done() // defer means always called in end of Function

	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}

	fmt.Println("Res for EndPoint: ", endpoint, " Res: ", res.StatusCode)
}
