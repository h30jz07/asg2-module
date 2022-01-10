package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Module struct {
	ModuleCode         string
	ModuleName         string
	Synopsis           string
	LearningObjectives string
	Classes            []string
	AssignedTutors     []string
	EnrolledStudents   []string
	RatingsAndComments []string
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Module API!")
}

func main() {
	os.Setenv("PORT", "8087")
	router := mux.NewRouter()
	router.HandleFunc("/module/v1/", health) //Health Check - database connectivity
	fmt.Println("Listening at port " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))
}

//Helper Functions
//Open Database Helper function
func OpenDatabase() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")

	//handle error
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened")
	}
	return db
}
