package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type ModuleDetails struct {
	ModuleCode         string `json:"modulecode"`
	ModuleName         string `json:"modulename"`
	Synopsis           string `json:"synopsis"`
	LearningObjectives string `json:"learningobjective"`
}

type ModuleDetails2 struct {
	ModuleCode         string `json:"modulecode"`
	ModuleName         string `json:"modulename"`
	Synopsis           string `json:"synopsis"`
	LearningObjectives string `json:"learningobjective"`
	TutorId            string `json:"tutorid"`
}

type EnrolledStudent struct {
	StudentId string `json:"student_id"`
	ClassId   int    `json:"class_id"`
	Semester  string `json:"semester"`
}

type AssignedTutor struct {
	TutorId    string `json:"tutorid"`
	ModuleCode string `json:"modulecode"`
}

type MoreDetails struct {
	EnrolledStudents []EnrolledStudent `json:"enrolled_students"`
	AssignedTutors   []AssignedTutor   `json:"assigned_tutors"`
	Classes          []int             `json:"classes"`
	RAndCLink        string            `json:"ratings_and_comments_url"`
}

type Module struct {
	ModuleCode         string            `json:"modulecode"`
	ModuleName         string            `json:"modulename"`
	Synopsis           string            `json:"synopsis"`
	LearningObjectives string            `json:"learningobjective"`
	Classes            []int             `json:"classes"`
	AssignedTutors     []AssignedTutor   `json:"assigned_tutors"`
	EnrolledStudents   []EnrolledStudent `json:"enrolled_students"`
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Module API!")
}

//List all modules
func listModules(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		modules := getAllModules()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(modules)
	}
}

func getAllModules() []ModuleDetails {
	os.Setenv("HOST_URL", "localhost")
	os.Setenv("MODULE_MANAGEMENT_PORT", "9141")
	//url := fmt.Sprintf(`http://%s:%s/api/v1/modules/`, os.Getenv("HOST_URL"), os.Getenv("MODULE_MANAGEMENT_PORT"))
	url := fmt.Sprintf(`http://%s:%s/modules`, os.Getenv("HOST_URL"), os.Getenv("MODULE_MANAGEMENT_PORT"))

	var modules []ModuleDetails
	if response, err := http.Get(url); err == nil {
		defer response.Body.Close()
		if body, _ := ioutil.ReadAll(response.Body); err == nil {
			json.Unmarshal(body, &modules)
			return modules
		}
	}
	return nil
}

//Get module details
func getModuleDetails(w http.ResponseWriter, r *http.Request) {
	os.Setenv("HOST_URL", "localhost")
	os.Setenv("R_AND_C_PORT", "9040")
	os.Setenv("TIMETABLE_PORT", "9141")
	os.Setenv("MODULE_MANAGEMENT_PORT", "9141")
	//Return link to 3.9(Ratings and Comments) and EnrolledStudents list
	switch r.Method {
	case "GET":
		//Retrieve moduleCode from path param
		moduleCode := mux.Vars(r)["moduleCode"]

		var moduleDetails MoreDetails

		//Get EnrolledStudents
		enrolledStudentsDetails := getEnrolledStudents(moduleCode)
		moduleDetails.EnrolledStudents = enrolledStudentsDetails
		//Get Classes
		moduleDetails.Classes = getClasses(enrolledStudentsDetails)
		//Get Assigned Tutors
		moduleDetails.AssignedTutors = getAssignedTutors(moduleCode)

		moduleDetails.RAndCLink = fmt.Sprintf("http://%s:%s/students?moduleCode=%s", os.Getenv("HOST_URL"), os.Getenv("R_AND_C_PORT"), moduleCode)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(moduleDetails)
	}
}

func getEnrolledStudents(moduleCode string) []EnrolledStudent {
	os.Setenv("HOST_URL", "localhost")
	os.Setenv("TIMETABLE_PORT", "9141")
	//url := fmt.Sprintf("http://%s:%s/api/v1/module/%s", os.Getenv("HOST_URL"), os.Getenv("TIMETABLE_PORT"), moduleCode) //Update for production
	url := fmt.Sprintf("http://%s:%s/%s", os.Getenv("HOST_URL"), os.Getenv("TIMETABLE_PORT"), moduleCode) //just for testing
	fmt.Println(url)
	if response, err := http.Get(url); err == nil {
		defer response.Body.Close()
		if body, _ := ioutil.ReadAll(response.Body); err == nil {
			var result []EnrolledStudent
			json.Unmarshal(body, &result)

			return result
		}
	}
	return nil
}

func getClasses(s []EnrolledStudent) []int {
	var classList []int
	for i := 0; i < len(s); i++ {
		if !intInSlice(s[i].ClassId, classList) {
			classList = append(classList, s[i].ClassId)
		}
	}
	return classList
}

func getAssignedTutors(moduleCode string) []AssignedTutor {
	os.Setenv("HOST_URL", "localhost")
	os.Setenv("MODULE_MANAGEMENT_PORT", "9141")
	//url := fmt.Sprintf("http://%s:%s/api/v1/module/tutor/%s", os.Getenv("HOST_URL"), os.Getenv("MODULE_MANAGEMENT_PORT"), moduleCode) //update for production
	url := fmt.Sprintf("http://%s:%s/tutor?modulecode=%s", os.Getenv("HOST_URL"), os.Getenv("MODULE_MANAGEMENT_PORT"), moduleCode) //just for testing
	if response, err := http.Get(url); err == nil {
		defer response.Body.Close()
		if body, _ := ioutil.ReadAll(response.Body); err == nil {
			var result []AssignedTutor
			json.Unmarshal(body, &result)

			return result
		}
	}
	return nil
}

func getModulesByTutor(w http.ResponseWriter, r *http.Request) {
	os.Setenv("HOST_URL", "localhost")
	os.Setenv("R_AND_C_PORT", "9040")
	os.Setenv("MODULE_MANAGEMENT_PORT", "9141")
	//Return link to 3.9(Ratings and Comments) and EnrolledStudents list
	switch r.Method {
	case "GET":
		//Retrieve tutorId from path param
		tutorId := mux.Vars(r)["tutorId"]

		//Get partial details of modules taught by tutor
		tutorModules := getModulesByTutorId(tutorId)

		var tutorModulesDetails []Module
		fmt.Println(tutorModules)

		for i := 0; i < len(tutorModules); i++ {
			enrolledStudentsDetails := getEnrolledStudents(tutorModules[i].ModuleCode)
			//Get modules details
			var moduleDetails Module
			moduleDetails.ModuleCode = tutorModules[i].ModuleCode
			moduleDetails.ModuleName = tutorModules[i].ModuleName
			moduleDetails.Synopsis = tutorModules[i].Synopsis
			moduleDetails.LearningObjectives = tutorModules[i].LearningObjectives
			moduleDetails.Classes = getClasses(enrolledStudentsDetails)
			moduleDetails.EnrolledStudents = enrolledStudentsDetails
			moduleDetails.AssignedTutors = getAssignedTutors(tutorModules[i].ModuleCode)

			tutorModulesDetails = append(tutorModulesDetails, moduleDetails)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tutorModulesDetails)
	}
}

func getModulesByTutorId(tutorId string) []ModuleDetails2 {
	os.Setenv("HOST_URL", "localhost")
	os.Setenv("MODULE_MANAGEMENT_PORT", "9141")
	//url := fmt.Sprintf("http://%s:%s/api/v1/module/tutor/%s", os.Getenv("HOST_URL"), os.Getenv("MODULE_MANAGEMENT_PORT"), moduleCode) //update for production
	url := fmt.Sprintf("http://%s:%s/tutor?tutorid=%s", os.Getenv("HOST_URL"), os.Getenv("MODULE_MANAGEMENT_PORT"), tutorId) //just for testing CHANGE WHEN AZZI UPDATE
	var result []ModuleDetails2
	if response, err := http.Get(url); err == nil {
		defer response.Body.Close()
		if body, _ := ioutil.ReadAll(response.Body); err == nil {
			json.Unmarshal(body, &result)
		}
	}

	return result
}

func main() {
	//environment variables
	//setup for local testing
	os.Setenv("PORT", "9061")
	os.Setenv("ORIGIN_ALLOWED", "*")

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router := mux.NewRouter()
	router.HandleFunc("/module/v1", health)                                //Health Check - database connectivity
	router.HandleFunc("/module/v1/list", listModules).Methods("GET")       //List Modules information
	router.HandleFunc("/module/v1/details/{moduleCode}", getModuleDetails) //get Modules information
	router.HandleFunc("/module/v1/modules/{tutorId}", getModulesByTutor)   //get Modules by tutorid
	fmt.Printf(`Listening at port %s`, os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

//Helper functions
func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
