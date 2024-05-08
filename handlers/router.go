package handlers

import (
	"employee-manager/store"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers routes for handling employee-related requests
func RegisterRoutes(router *mux.Router, store store.EmployeeStore) {
	employeeHandler := newEmployeeHandler(store)

	router.HandleFunc("/v1/employees", employeeHandler.listEmployees).Methods("GET")
	router.HandleFunc("/v1/employees", employeeHandler.createEmployee).Methods("POST")
	router.HandleFunc("/v1/employees/{id}", employeeHandler.getEmployeeByID).Methods("GET")
	router.HandleFunc("/v1/employees/{id}", employeeHandler.updateEmployee).Methods("PUT")
	router.HandleFunc("/v1/employees/{id}", employeeHandler.deleteEmployee).Methods("DELETE")
}
