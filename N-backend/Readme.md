# Concept Of Go Lang net/http

1. Minimal Http Code
   - It's Interface which has Method `ServeHTTP(w http.ResponseWriter, r *http.Request)`
   - `HandlerFunc` Type: Adapter to convert a function into a Handler.
   - `HandleFunc(string , func(ResponseWriter, Request))`

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

// Minimal Hello World Http Server
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	// on /hello run the HelloHandler
	http.HandleFunc("/hello", HelloHandler)
	log.Println("Running Http Server on :8080")

	// Server Listening on 8080
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Error Running Http Server on :8080")
		return
	}
}
```

2. http.Handler

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

// Minimal Hello World Http Server
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

type myHandler struct{} // empty struct

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi from custom handler")
}

func main() {

	// its multiplexer
	mux := http.NewServeMux()

	// takes func(ResponseWriter , Request)
	mux.HandleFunc("/hello", HelloHandler)

	// HandlerFunc(func(ResponseWriter,Request))
	// Converts HelloHandler to http.Handler
	mux.Handle("/handler", http.HandlerFunc(HelloHandler))

	// myHandler is custom http.Handler which is executing for route /myhandler
	mux.Handle("/myhandler", &myHandler{}) // it supports all the routes

	log.Print("Server Running on 8080")

	// mux also implements ServerHttp thats why its valid
	http.ListenAndServe(":8080", mux)
}
```

3. Middleware

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request Receiver: ", r.Method, " ", r.URL.Path)

		// internally MiddlewareRoute -> HandlerFunc which implements ServerHttp
		// inside that this below code is there
		// and f means MiddlewareRoute here
		// and directly f(w,r) means MiddlewareRoute(w,r)
		/*
			func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
				f(w, r)
			}
		*/
		next.ServeHTTP(w, r)
	})
}

func MiddlewareRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello this is Custom Logging Middleware")
}

func main() {
	http.Handle("/middleware", LoggingMiddleware(http.HandlerFunc(MiddlewareRoute)))

	log.Println("Server Running on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
```

3. Internal Handler

```go
package main

import "fmt"

// Interface defining one method
type Handler interface {
	serverHttp()
}

// Function type
type HandlerFunc func()

// Method implemented on the function type
func (h HandlerFunc) serverHttp() {
	fmt.Println("Inside serverHttp method")
	h() // calling the actual function
}

// Function returning Handler interface
func something() Handler {
	// Assign an actual function to HandlerFunc variable
	var a HandlerFunc = func() {
		fmt.Println("Actual HandlerFunc body running")
	}

	fmt.Println("Works")
	return a // 'a' satisfies Handler interface because it has serverHttp()
}

func main() {
	a := something()
	a.serverHttp() // calling interface method
}
```
