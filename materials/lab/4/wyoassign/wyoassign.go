package wyoassign

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"

)

type Response struct{
	Assignments []Assignment `json:"assignments"`
}
type Responses struct{
	Courses []Course `json:"courses"`
}

type Assignment struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"desc"`
	Points int `json:"points"`
	//DueDate string `json:"duedate"`
}
type Course struct{
	Id string `json:"id"`
	Course string `json:"course"` 
	CourseNumber int `json:"coursenumber"`
}

var Assignments []Assignment
const Valkey string = "FooKey"

var Courses []Course

func InitAssignments(){
	var assignmnet Assignment
	assignmnet.Id = "Mike1A"
	assignmnet.Title = "Lab 4 "
	assignmnet.Description = "Some lab this guy made yesteday?"
	assignmnet.Points = 20
	//assignmnet.DueDate = "3/11/2022";
	Assignments = append(Assignments, assignmnet)
}
func InitCourses(){
	var cours Course
	cours.Id = "Tuesdays"
	cours.Course = "Linux"
	cours.CourseNumber = 3750
	//assignmnet.DueDate = "3/11/2022";
	Courses = append(Courses, cours)
}

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}


func GetAssignments(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Assignments = Assignments

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	//TODO 
	w.Write(jsonResponse)
}
func GetCourses(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Responses

	response.Courses = Courses

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	//TODO 
	w.Write(jsonResponse)
}

func GetAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	
	params := mux.Vars(r)

	for _, assignment := range Assignments {
		if assignment.Id == params["id"]{
			json.NewEncoder(w).Encode(assignment)
			w.WriteHeader(http.StatusOK)
			break
		}else{
			w.WriteHeader(http.StatusNotFound)
		}
	}
	//TODO : Provide a response if there is no such assignment
	//w.Write(jsonResponse)
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	
	params := mux.Vars(r)

	for _, course := range Courses {
		if course.Id == params["id"]{
			json.NewEncoder(w).Encode(course)
			w.WriteHeader(http.StatusOK)
			break
		}else{
			w.WriteHeader(http.StatusNotFound)
		}
	}
	//TODO : Provide a response if there is no such assignment
	//w.Write(jsonResponse)
}

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	
	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, assignment := range Assignments {
			if assignment.Id == params["id"]{
				Assignments = append(Assignments[:index], Assignments[index+1:]...)
				response["status"] = "Success"
				break
			}
	}
		
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	
	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, course := range Courses {
			if course.Id == params["id"]{
				Courses = append(Courses[:index], Courses[index+1:]...)
				response["status"] = "Success"
				break
			}
	}
		
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	
	var response Response
	response.Assignments = Assignments

	params := mux.Vars(r)

	var assignmnet Assignment
	r.ParseForm()
	for index, assignment := range Assignments {
		if assignment.Id == params["id"]{
			Assignments = append(Assignments[:index], Assignments[index+1:]...)
			assignmnet.Id =  r.FormValue("id")
			assignmnet.Title =  r.FormValue("title")
			assignmnet.Description =  r.FormValue("desc")
			assignmnet.Points, _ =  strconv.Atoi(r.FormValue("points"))
			//assignmnet.DueDate, _ = r.FormValue("duedate")
			Assignments = append(Assignments, assignmnet)
			w.WriteHeader(http.StatusCreated)
			break
		}else{
			w.WriteHeader(http.StatusNotFound)
		}
		
		
	}
	//w.Write(jsonResponse)
}
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	
	var response Responses
	response.Courses = Courses

	params := mux.Vars(r)

	var cours Course
	r.ParseForm()
	for index, course := range Courses {
		if course.Id == params["id"]{
			Courses = append(Courses[:index], Courses[index+1:]...)
			cours.Id = r.FormValue("id")
			cours.Course = r.FormValue("course")
			cours.CourseNumber, _ = strconv.Atoi(r.FormValue("coursenumber"))
			Courses = append(Courses, cours)
			w.WriteHeader(http.StatusCreated)
			break
		}else{
			w.WriteHeader(http.StatusNotFound)
		}
		
		
	}
	//w.Write(jsonResponse)
}
func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var assignmnet Assignment
	r.ParseForm()
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging
	if(r.FormValue("id") != ""){
		
		assignmnet.Id =  r.FormValue("id")
		assignmnet.Title =  r.FormValue("title")
		assignmnet.Description =  r.FormValue("desc")
		assignmnet.Points, _ =  strconv.Atoi(r.FormValue("points"))
		//assignmnet.DueDate, _ = r.FormValue("duedate")
		Assignments = append(Assignments, assignmnet)
		w.WriteHeader(http.StatusCreated)
	}else{
		w.WriteHeader(http.StatusNotFound)
	}

}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cours Course
	r.ParseForm()
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging
	if(r.FormValue("id") != ""){
		
		cours.Id = r.FormValue("id")
		cours.Course = r.FormValue("course")
		cours.CourseNumber, _ = strconv.Atoi(r.FormValue("coursenumber"))
		Courses = append(Courses, cours)
		w.WriteHeader(http.StatusCreated)
	}else{
		w.WriteHeader(http.StatusNotFound)
	}

}