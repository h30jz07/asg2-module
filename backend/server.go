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

type ModuleDetailsTutor struct {
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
	TutorId     int    `json:"tutor_id"`
	TutorName   string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"descriptions"`
	ModuleCode  string `json:"modulecode"`
	ModuleId    string `json:"moduleid"`
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		modules := getAllModules()
		json.NewEncoder(w).Encode(modules)
	}
}

func getAllModules() []ModuleDetails {
	url := fmt.Sprintf(`http://%s:%s/api/v1/modules/`, os.Getenv("HOST_URL"), os.Getenv("MODULE_MANAGEMENT_PORT"))

	var modules []ModuleDetails
	if response, err := http.Get(url); err == nil {
		defer response.Body.Close()
		if body, _ := ioutil.ReadAll(response.Body); err == nil {
			json.Unmarshal(body, &modules)
			return modules
		}
	} else {
		fmt.Println(err)
	}
	return nil
}

//Get module details
func getModuleDetails(w http.ResponseWriter, r *http.Request) {
	os.Setenv("HOST_URL", "localhost")
	os.Setenv("R_AND_C_PORT", "9040")
	switch r.Method {
	case "GET":
		w.Header().Set("Access-Control-Allow-Origin", "*")
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
		fmt.Println(moduleDetails)

		var moduleId string
		if moduleDetails.AssignedTutors != nil {
			moduleId = moduleDetails.AssignedTutors[0].ModuleId
		} else {

			moduleId = "0"
		}

		moduleDetails.RAndCLink = fmt.Sprintf("http://%s:%s/Main/details.html?id=%stype=Module", os.Getenv("HOST_URL"), os.Getenv("R_AND_C_PORT"), moduleId)
		json.NewEncoder(w).Encode(moduleDetails)
	}
}

func getEnrolledStudents(moduleCode string) []EnrolledStudent {
	url := fmt.Sprintf("http://%s:%s/api/v1/allocations/module/%s", os.Getenv("HOST_URL"), os.Getenv("TIMETABLE_PORT"), moduleCode) //Update for production
	fmt.Println("Calling Get Enrolled Students from Timetable MS")
	if response, err := http.Get(url); err == nil {
		defer response.Body.Close()
		if body, _ := ioutil.ReadAll(response.Body); err == nil {
			var result []EnrolledStudent
			json.Unmarshal(body, &result)

			return result
		}
	} else {
		fmt.Println(err)
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
	url := fmt.Sprintf("http://%s:%s/api/v1/module/tutor/%s", os.Getenv("HOST_URL"), os.Getenv("MODULE_MANAGEMENT_PORT"), moduleCode) //update for production
	fmt.Println("Calling Get Assigned Tutors from Module Management MS")
	if response, err := http.Get(url); err == nil {
		defer response.Body.Close()
		if body, _ := ioutil.ReadAll(response.Body); err == nil {
			var result []AssignedTutor
			json.Unmarshal(body, &result)

			return result
		}
	} else {
		fmt.Println(err)
	}
	return nil
}

func getModulesByTutor(w http.ResponseWriter, r *http.Request) {
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(tutorModulesDetails)
	}
}

func getModulesByTutorId(tutorId string) []ModuleDetailsTutor {
	url := fmt.Sprintf("http://%s:%s/api/v1/module/alltutor/%s", os.Getenv("HOST_URL"), os.Getenv("MODULE_MANAGEMENT_PORT"), tutorId) //update for production
	var result []ModuleDetailsTutor
	if response, err := http.Get(url); err == nil {
		defer response.Body.Close()
		if body, _ := ioutil.ReadAll(response.Body); err == nil {
			json.Unmarshal(body, &result)
		}
	} else {
		fmt.Println(err)
	}

	return result
}

func main() {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router := mux.NewRouter()
	router.HandleFunc("/module/v1", health)                                                          //Health Check - database connectivity
	router.HandleFunc("/module/v1/list", listModules).Methods("GET", "OPTIONS")                      //List Modules information
	router.HandleFunc("/module/v1/details/{moduleCode}", getModuleDetails).Methods("GET", "OPTIONS") //get Modules information
	router.HandleFunc("/module/v1/modules/{tutorId}", getModulesByTutor).Methods("GET", "OPTIONS")   //get Modules by tutorid
	fmt.Printf(`Listening at port %s`, os.Getenv("BACKEND_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("BACKEND_PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(router)))
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
