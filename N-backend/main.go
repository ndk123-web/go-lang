package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserData struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Email    string   `json:"email"`
	Password string   `json:"-"`              // it means dont include password in json
	Tags     []string `json:"tags,omitempty"` // it means if nil then dont add in json
}

func main() {
	fmt.Println("Sending And Accepting Req in Go Lang")

	// HandleFunc like app.get(response , request)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World bc")
	})

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"hello": "Hi I Am ndk"})
	})

	// sends json data
	// struct -> Encode (struct->JSON) -> Byte Class -> Binary Stream (to Browser)
	http.HandleFunc("/user/ndk", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		user := []UserData{
			{"Navnath", 20, "ndk@gmail.com", "pass1", []string{"Hi", "Bro"}},
			{"Harsh", 20, "hasrsh@gmail.com", "pass12", nil},
		}

		jsonData, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(jsonData))

		// it sends actual json Response
		// NewEncoder(w) accept io.Writer where w is ResponseWriter which implement interface io.Writer
		// Struct → JSON → Byte slice → ResponseWriter → TCP → Browser
		json.NewEncoder(w).Encode(user)
	})

	// receiver (Long)
	// receiver json data and modify and send
	// receiver -> Binary Stream -> get []byte body -> JSON -> struct
	http.HandleFunc("/receive/u", func(w http.ResponseWriter, r *http.Request) {

		rawBody, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		var user UserData

		err := json.Unmarshal(rawBody, &user)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Received user: %+v\n", user)

		// modify if needed and send back
		user.Name = user.Name + "_modified"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	// receiver (short)
	// receiver json data and modify and send
	// receiver -> Binary Stream -> get []byte body -> JSON -> struct
	http.HandleFunc("/receive/u2", func(w http.ResponseWriter, r *http.Request) {

		var user UserData
		json.NewDecoder(r.Body).Decode(&user)

		w.Header().Set("Content-Type", "application/json")

		user.Name = user.Name + "_modified"

		json.NewEncoder(w).Encode(&user)
	})

	fmt.Println("Server Started on Port 3000")

	DecodeJson()
	http.ListenAndServe(":3000", nil)
}

func DecodeJson() {
	jsonData := []byte(`
		{
			"name":"ndk",
			"age": 20,
			"email":"ndk@gmail.com",
			"password": "pass123",
			"tags": ["op","bro"]
		}
	`)

	var user UserData

	checkIsValid := json.Valid(jsonData)
	if checkIsValid {
		fmt.Println("JSON Was Valid")
		err := json.Unmarshal(jsonData, &user)
		if err != nil {
			panic(err)
		}
		fmt.Println("User: ", user)
		fmt.Print("Pass: ", user.Password)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}
}
