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

//we will create a learn code online with backed, we will work with the courses
//user can get all the courses, user can create new courses and update the exiting courses
//if their is no unique id no tilte or the name of the course then course is not edit
//so we need to create a helper function for that so we will use slice for fake data base

//after all that we will inject gorilla mux so that every single route is handle by
//a seperate method for all th e curd(create update read delete)operation

// Model for courses - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json: "author"`
}
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"Website"`
}

// fake database
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("API - LearnCodeConline.in")
	r :=mux.NewRouter()


	//seeding
	courses=append(courses, Course{CourseId: "2",CourseName: "ReactJS",CoursePrice: 299,Author: &Author{FullName:
		 "Hitesh Choudhary",Website: "lco.dev"}})
	courses=append(courses, Course{CourseId: "4",CourseName: "MERN Stack",CoursePrice: 199,Author: &Author{FullName:
			"Hitesh Choudhary",Website: "go.dev"}})	
	//routing......(/) is route
	r.HandlerFunc("/",serveHome).Methods("GET")
	r.HandlerFunc("/courses",getAllCourses).Methods("GET")
	r.HandlerFunc("/course/{id}",getOnecourse).Methods("GET")
	r.HandlerFunc("/courses",createOneCourse).Methods("POST")
	r.HandlerFunc("/courses/{id}",updateOneCourse).Methods("PUT")
	



	//listen to a port
	log.Fatal(http.ListenAndServe(":4000",r))		
    
}

//controllers -file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCoderline</h1>"))
}

// another rotuer goal is through up all values from my database
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All courses")
	//set headers
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOnecourse(w http.ResponseWriter, r *http.Request) {
	//grab the unique id
	//since we all courses as an array just go and loop through that array
	//and compare this id with array that id is present or not
	fmt.Println("Get one array")
	w.Header().Set("Content-Type", "application/json")

	//need to gran id from request
	params := mux.Vars(r)
	//loop through the courses,find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found of this id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what if :body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about -{}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return

	}

	//generate unique id,string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Int())
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter,r http.Request){
	fmt.Println("Update one course")
	w.Header().Set("Content-Type","application/json")

	//first -grab id from req
	params:=mux.Vars(r)
	//lops,id,remove,add with my ID

	for index, course :=range courses{
		if course.CourseId==params["id"]{
			courses=append(courses[:index],courses[index+1:]...)
			var course Course
			_=json.NewDecoder(r.Body).Decode(&course)
			course.CourseId=params["id"]
			courses=append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//todo send a response when id is not found
}
