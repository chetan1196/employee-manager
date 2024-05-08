package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"employee-manager/employee"
	"employee-manager/store"

	"github.com/gorilla/mux"
)

const (
	// Default page and pageSize parameters
	defaultPage     = 1
	defaultPageSize = 10
)

// EmployeeHandler handles HTTP requests related to employees
type EmployeeHandler struct {
	store store.EmployeeStore
}

// NewEmployeeHandler creates a new instance of EmployeeHandler
func newEmployeeHandler(store store.EmployeeStore) *EmployeeHandler {
	return &EmployeeHandler{
		store: store,
	}
}

// ListEmployees handles listing employees
func (h *EmployeeHandler) listEmployees(w http.ResponseWriter, r *http.Request) {
	var (
		page, pageSize = defaultPage, defaultPageSize
		err            error
	)
	pageStr := r.URL.Query().Get("page")
	if len(pageStr) > 0 {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			errMsg := APIError{Message: "Invalid page number"}
			writeError(w, errMsg, http.StatusBadRequest)
		}
	}

	pageSizeStr := r.URL.Query().Get("pageSize")
	if len(pageSizeStr) > 0 {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil {
			errMsg := APIError{Message: "Invalid page size"}
			writeError(w, errMsg, http.StatusBadRequest)
		}
	}

	log.Printf("Page: %d, PageSize: %d", page, pageSize)

	employees := h.store.List(page, pageSize)

	log.Printf("Number of employees: %d", len(employees))

	writeSuccess(w, employees, http.StatusOK)
}

// CreateEmployee handles creating a new employee
func (h *EmployeeHandler) createEmployee(w http.ResponseWriter, r *http.Request) {
	var emp employee.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		errMsg := APIError{Message: "Invalid request body"}
		writeError(w, errMsg, http.StatusBadRequest)
		return
	}

	id := h.store.GenerateID()

	emp.ID = id

	h.store.Create(emp)

	resp := map[string]interface{}{
		"id": id,
	}

	writeSuccess(w, resp, http.StatusCreated)
}

// GetEmployeeByID handles retrieving an employee by ID
func (h *EmployeeHandler) getEmployeeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errMsg := APIError{Message: "Invalid employee ID"}
		writeError(w, errMsg, http.StatusBadRequest)
		return
	}

	emp, ok := h.store.GetByID(id)
	if !ok {
		errMsg := APIError{Message: "Employee does not exists"}
		writeError(w, errMsg, http.StatusNotFound)
		return
	}

	writeSuccess(w, emp, http.StatusOK)
}

// UpdateEmployee handles updating an employee
func (h *EmployeeHandler) updateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errMsg := APIError{Message: "Invalid employee ID"}
		writeError(w, errMsg, http.StatusBadRequest)
		return
	}

	var newEmp employee.Employee
	err = json.NewDecoder(r.Body).Decode(&newEmp)
	if err != nil {
		errMsg := APIError{Message: "Invalid request body"}
		writeError(w, errMsg, http.StatusBadRequest)
		return
	}

	if !h.store.Update(id, newEmp) {
		errMsg := APIError{Message: "Employee does not exists"}
		writeError(w, errMsg, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteEmployee handles deleting an employee
func (h *EmployeeHandler) deleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errMsg := APIError{Message: "Invalid employee ID"}
		writeError(w, errMsg, http.StatusBadRequest)
		return
	}

	if !h.store.Delete(id) {
		errMsg := APIError{Message: "Employee does not exists"}
		writeError(w, errMsg, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
