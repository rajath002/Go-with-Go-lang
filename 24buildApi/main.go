package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullName"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper -file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", createOneCourse).Methods("POST")
	// r.HandleFunc("/", updateOneCourse).Methods("PUT")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers -file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to API by Rajath</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	// Grab id from the request
	params := mux.Vars(r)

	fmt.Println("Params ", params)

	//loop through courses and find the matching Id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course find with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about  - {}
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No Data inside JSON")
		return
	}

	// Generate unique ID
	// Append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Course")
	w.Header().Set("Content-Type", "application/json")

	// First - grab id from req
	params := mux.Vars(r)

	//  loop, id, remove, add with my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// TODO: Send a response when id is not found!

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// loop, id, remove (index, index+1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode("Course deleted successfully!")
}
