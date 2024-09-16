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

// models for course
type Course struct {
	Id     string  `json:"courseid`
	Name   string  `json:"courseName"`
	Price  int     `json:"price"`
	Author *Author `json:"author"`
}
type Author struct {
	Name string `json:"authorName`
	Link string `json:"link"`
}

//fake DB

var courses []Course

// middleware

func (c *Course) IsEmpty() bool {
	return c.Name == ""
}

// main app
func crudRestApi() {
	fmt.Println("lets go")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{Id: "1", Name: "Go Lang", Price: 1999, Author: &Author{Name: "shubham", Link: "theshubham.dev"}})
	courses = append(courses, Course{Id: "2", Name: "ReactJs", Price: 199, Author: &Author{Name: "shubham", Link: "theshubham.dev"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCoursebyId).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4001", r))

}

// controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welome to go lang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getCoursebyId(w http.ResponseWriter, r *http.Request) {
	// input form request
	params := mux.Vars(r)
	fmt.Println("Course Id:", params["id"])
	w.Header().Set("Content-Type", "application/json")
	// find course with id and return as response

	for _, course := range courses {
		if course.Id == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found")
	return

}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Course Created")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Validation error")
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Not Valid data as Paylaod")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.Id = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// take input form request params and payload
	params := mux.Vars(r)
	fmt.Println("Course Id:", params["id"])

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Not Valid data as Paylaod")
		return
	}

	for index, course := range courses {
		if course.Id == params["id"] {
			// remove data using id
			courses = append(courses[:index], courses[index+1:]...)
			// adding new data
			course.Id = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found")
	return

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// take input form request params and payload
	params := mux.Vars(r)
	fmt.Println("Course Id:", params["id"])

	for index, course := range courses {
		if course.Id == params["id"] {
			// remove data using id
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("course deleted Succesfully")
			break
		}
	}
}
