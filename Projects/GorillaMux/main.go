package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Models for courses - file
type Course struct {
	CourseId   string  `json:"courseId"`
	CourseName string  `json:"courseName"`
	Price      float64 `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// DB - file
var Courses = []Course{
	{"1", "AIML", 12.24, &Author{"NDK", "https://ndkdev.me"}},
	{"2", "JS", 11.24, &Author{"HARS", "https://harsh.me"}},
}

// Middleware / Helpers - file
func (c *Course) CourseIsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

// Controller-file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello I Am Navnath</h1>"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Courses)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")

	w.Header().Set("Content-Type", "application/json")

	// Get the Varabile (Params /{id})
	params := mux.Vars(r)
	fmt.Println("Params: ", params)

	// fetch the id from that params
	id := params["id"]
	found := false
	var FoundCourse Course

	// Loop through over all Courses
	for _, v := range Courses {
		if v.CourseId == id {
			found = true
			FoundCourse = v
			break
		}
	}

	if found {
		// Encode can Take Any Of Type
		json.NewEncoder(w).Encode(FoundCourse)
	} else {
		// Encode can Take Any Of Type
		json.NewEncoder(w).Encode(map[string]string{"Error": "Not Found The Given Course For Give Id"})
	}
}

func AddCourseController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// what if body empty ?
	if r.Body == nil {
		json.NewEncoder(w).Encode(map[string]string{"Error": "Body is Empty"})
		return
	}

	// Decode and save into the decodeJsonData
	var decodedJsonData Course
	_ = json.NewDecoder(r.Body).Decode(&decodedJsonData)

	// It Checks Course Id and Name is Empty or not
	if decodedJsonData.CourseIsEmpty() {
		json.NewEncoder(w).Encode(map[string]string{"Error": "CoursIsEmpty"})
		return
	}

	// append in Course again
	Courses = append(Courses, decodedJsonData)

	fmt.Println("After Adding Course")
	for _, v := range Courses {
		fmt.Println(v)
	}

	json.NewEncoder(w).Encode(map[string]string{"Success": "True"})
}

func UpdateController(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		json.NewEncoder(w).Encode(map[string]string{"Error": "Body is Empty"})
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var requestData Course

	// get the params
	params := mux.Vars(r)

	// get the requested updated data
	_ = json.NewDecoder(r.Body).Decode(&requestData)

	// check is id is there
	for i, v := range Courses {
		if v.CourseId == params["id"] {
			Courses[i].CourseName = requestData.CourseName
			Courses[i].Author = requestData.Author
			Courses[i].Price = requestData.Price
			break
		}
	}

	fmt.Println("After Updating Course")
	for _, v := range Courses {
		fmt.Println(v)
	}

	json.NewEncoder(w).Encode(map[string]string{"Success": "Updated"})
}

func DeleteCourseController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Found := false

	// if id is there then no issue
	id, ok := params["id"]
	if !ok {
		json.NewEncoder(w).Encode(map[string]string{"Error": "For Delete Id is Not there"})
		return
	}

	for i, v := range Courses {
		if v.CourseId == id {

			// delete that its like substring
			Courses = append(Courses[:i], Courses[i+1:]...)
			Found = true
			break
		}
	}

	// Its like Not found Course Id
	if !Found {
		json.NewEncoder(w).Encode(map[string]string{"Error": "Course ID Not Found"})
		return
	}

	// If found and all go as planned
	fmt.Println("After Deleting Course")
	for _, v := range Courses {
		fmt.Println(v)
	}

	json.NewEncoder(w).Encode(map[string]string{"Success": "Deleted"})
}

func main() {
	fmt.Println("Mux in Go Lang")

	// create a router
	r := mux.NewRouter()

	// create handler function for GET Request
	r.HandleFunc("/", serveHome).Methods("GET")

	// get all courses
	r.HandleFunc("/courses", GetAllCourses).Methods("GET")

	// get one course
	r.HandleFunc("/course/{id}", GetOneCourse).Methods("GET")

	// add course
	r.HandleFunc("/course/add", AddCourseController).Methods("POST")

	// update course
	r.HandleFunc("/course/update/{id}", UpdateController).Methods("POST")

	// delete course
	r.HandleFunc("/course/delete/{id}", DeleteCourseController).Methods("POST")

	// run the server and give the reader
	http.ListenAndServe(":3000", r)
}
