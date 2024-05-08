package handlers

import (
	"bytes"
	"employee-manager/employee"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// MockEmployeeStore is a mock implementation of the EmployeeStore interface
type MockEmployeeStore struct {
	employees map[int]employee.Employee
	idCounter int
}

func (m *MockEmployeeStore) GenerateID() int {
	m.idCounter++
	return m.idCounter
}

func (m *MockEmployeeStore) Create(emp employee.Employee) {
	m.employees[emp.ID] = emp
}

func (m *MockEmployeeStore) GetByID(id int) (employee.Employee, bool) {
	emp, ok := m.employees[id]
	return emp, ok
}

func (m *MockEmployeeStore) Update(id int, newEmp employee.Employee) bool {
	if _, ok := m.employees[id]; !ok {
		return false
	}
	m.employees[id] = newEmp
	return true
}

func (m *MockEmployeeStore) Delete(id int) bool {
	if _, ok := m.employees[id]; !ok {
		return false
	}
	delete(m.employees, id)
	return true
}

func (m *MockEmployeeStore) List(page, pageSize int) []employee.Employee {

	start := (page - 1) * pageSize
	end := start + pageSize
	if end > len(m.employees) {
		end = len(m.employees)
	}

	employees := make([]employee.Employee, 0, pageSize)
	i := 0
	for _, emp := range m.employees {
		if i >= start && i < end {
			employees = append(employees, emp)
		}
		i++
	}

	return employees
}

// TestCases
func TestListEmployees(t *testing.T) {
	mockStore := &MockEmployeeStore{
		employees: map[int]employee.Employee{
			1: {ID: 1, Name: "Chetan Tiwari", Position: "Software Engineer", Salary: 50000},
			2: {ID: 2, Name: "Prince Sharma", Position: "Product Manager", Salary: 60000},
			3: {ID: 3, Name: "Vijay Pandey", Position: "Data Scientist", Salary: 70000},
		},
	}

	employeeHandler := newEmployeeHandler(mockStore)

	req, err := http.NewRequest("GET", "/v1/employees?page=1&pageSize=2", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	employeeHandler.listEmployees(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var employees []employee.Employee
	err = json.NewDecoder(recorder.Body).Decode(&employees)
	assert.NoError(t, err)

	assert.Len(t, employees, 2)
}

func TestCreateEmployee(t *testing.T) {
	mockStore := &MockEmployeeStore{
		employees: make(map[int]employee.Employee),
	}

	employeeHandler := newEmployeeHandler(mockStore)

	emp := employee.Employee{ID: 1, Name: "Chetan Tiwari", Position: "Software Engineer", Salary: 50000}
	body, err := json.Marshal(emp)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/v1/employees", bytes.NewReader(body))
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	employeeHandler.createEmployee(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	_, ok := mockStore.GetByID(emp.ID)
	assert.True(t, ok)
}

func TestGetEmployeeByID(t *testing.T) {
	mockStore := &MockEmployeeStore{
		employees: map[int]employee.Employee{
			1: {ID: 1, Name: "Chetan Tiwari", Position: "Software Engineer", Salary: 50000},
		},
	}

	req, err := http.NewRequest("GET", "/v1/employees/1", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	employeeHandler := newEmployeeHandler(mockStore)

	// Serve HTTP request using Gorilla Mux router
	router := mux.NewRouter()
	router.HandleFunc("/v1/employees/{id}", employeeHandler.getEmployeeByID).Methods("GET")
	router.ServeHTTP(recorder, req)

	employeeHandler.getEmployeeByID(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var emp employee.Employee
	err = json.NewDecoder(recorder.Body).Decode(&emp)
	assert.NoError(t, err)

	expectedEmp, _ := mockStore.GetByID(1)
	assert.Equal(t, expectedEmp.Name, emp.Name)
	assert.Equal(t, expectedEmp.Position, emp.Position)
	assert.Equal(t, expectedEmp.Salary, emp.Salary)

}

func TestUpdateEmployee(t *testing.T) {
	mockStore := &MockEmployeeStore{
		employees: map[int]employee.Employee{
			1: {ID: 1, Name: "Chetan Tiwari", Position: "Software Engineer", Salary: 50000},
		},
	}

	newEmp := employee.Employee{ID: 1, Name: "Chetan Tiwari", Position: "Senior Engineer", Salary: 60000}
	body, err := json.Marshal(newEmp)
	assert.NoError(t, err)

	req, err := http.NewRequest("PUT", "/v1/employees/1", bytes.NewReader(body))
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	employeeHandler := newEmployeeHandler(mockStore)

	// Serve HTTP request using Gorilla Mux router
	router := mux.NewRouter()
	router.HandleFunc("/v1/employees/{id}", employeeHandler.updateEmployee).Methods("PUT")
	router.ServeHTTP(recorder, req)

	employeeHandler.updateEmployee(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	updatedEmp, _ := mockStore.GetByID(1)
	assert.Equal(t, updatedEmp.Name, newEmp.Name)
	assert.Equal(t, updatedEmp.Position, newEmp.Position)
	assert.Equal(t, updatedEmp.Salary, newEmp.Salary)
}

func TestDeleteEmployee(t *testing.T) {
	mockStore := &MockEmployeeStore{
		employees: map[int]employee.Employee{
			1: {ID: 1, Name: "Chetan Tiwari", Position: "Software Engineer", Salary: 50000},
		},
	}

	req, err := http.NewRequest("DELETE", "/v1/employees/1", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	employeeHandler := newEmployeeHandler(mockStore)

	// Serve HTTP request using Gorilla Mux router
	router := mux.NewRouter()
	router.HandleFunc("/v1/employees/{id}", employeeHandler.deleteEmployee).Methods("DELETE")
	router.ServeHTTP(recorder, req)

	employeeHandler.deleteEmployee(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	_, ok := mockStore.GetByID(1)
	assert.False(t, ok)
}
