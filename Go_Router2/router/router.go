package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type employee struct {
	ID        int     `json:ID`
	Name      string  `json:Name`
	Basic_Pay float64 `json:Basic Pay`
	Dept_Name string  `json:Dept Name`
	HRA       float64
	GrossPay  float64
}

/*
Request by client (POST HTTP )
id
Name
Basic Pay   50000
Dept_Name   IT/Security
HRA         IT =10%BP  Security 14%BP
Gross Pay       0
------------------
Response by server
id
Name
Basic Pay   50000
HRA         10000
Gross Pay   60000
*/

type allEmployees []employee // list of employees

var employees = allEmployees{
	{
		ID:        1,
		Name:      "Employee One",
		Basic_Pay: 50000,
		Dept_Name: "IT/Security",
		HRA:       10000,
		GrossPay:  60000,
	},
}

//HTTP GET
func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	var newEmployee employee
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(reqBody, &newEmployee)

	newEmployee.ID = len(employees) + 1

	switch newEmployee.Dept_Name {
	case "IT":
		newEmployee.HRA = newEmployee.Basic_Pay * 0.1
	case "Security":
		newEmployee.HRA = newEmployee.Basic_Pay * 0.14
	default:
		newEmployee.HRA = 0
	}

	newEmployee.GrossPay = newEmployee.Basic_Pay + newEmployee.HRA
	employees = append(employees, newEmployee)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEmployee)

}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, emp := range employees {
		if emp.ID == empID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(emp)
		}
	}
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for index, emp := range employees {
		if emp.ID == empID {
			employees = append(employees[:index], employees[index+1:]...)
			fmt.Fprintf(w, "Employee with ID %v has been removed successfully", empID)
		}
	}
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empID, err := strconv.Atoi(vars["id"])
	var updateEmp employee

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}
	json.Unmarshal(reqBody, &updateEmp)

	switch updateEmp.Dept_Name {
	case "IT":
		updateEmp.HRA = updateEmp.Basic_Pay * 0.1
	case "Security":
		updateEmp.HRA = updateEmp.Basic_Pay * 0.14
	default:
		updateEmp.HRA = 0
	}

	updateEmp.GrossPay = updateEmp.Basic_Pay + updateEmp.HRA

	for index, ee := range employees {
		if ee.ID == empID {
			employees = append(employees[:index], employees[index+1:]...)
			updateEmp.ID = empID
			employees = append(employees, updateEmp)

			fmt.Fprintf(w, "The employee with id %v has been updated successfully", empID)
		}
	}
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}

//ConfigureRouter setup the router
func ConfigureRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/employees", getEmployees).Methods("GET")
	router.HandleFunc("/employees", createEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}", getEmployee).Methods("GET")
	router.HandleFunc("/employees/{id}", deleteEmployee).Methods("DELETE")
	router.HandleFunc("/employees/{id}", updateEmployee).Methods("PUT")
	router.HandleFunc("/ping", healthCheck).Methods("GET")
	router.Use(mux.CORSMethodMiddleware(router))

	return router
}
