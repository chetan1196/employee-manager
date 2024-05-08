package store

import (
	"employee-manager/employee"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateInMemoryStore(t *testing.T) {
	store := NewInMemoryStore()
	emp := employee.Employee{ID: 1, Name: "Chetan Tiwari", Position: "Software Engineer", Salary: 50000}

	store.Create(emp)

	result, ok := store.GetByID(emp.ID)
	assert.True(t, ok)
	assert.Equal(t, emp, result)
}

func TestGetByIdEmployee(t *testing.T) {
	store := NewInMemoryStore()
	emp := employee.Employee{ID: 1, Name: "Chetan Tiwari", Position: "Software Engineer", Salary: 50000}

	store.Create(emp)

	result, ok := store.GetByID(emp.ID)
	assert.True(t, ok)
	assert.Equal(t, emp, result)
}

func TestUpdateEmployee(t *testing.T) {
	store := NewInMemoryStore()
	emp := employee.Employee{ID: 1, Name: "Chetan Tiwari", Position: "Software Engineer", Salary: 50000}
	newEmp := employee.Employee{ID: 1, Name: "Chetan Tiwari", Position: "Senior Engineer", Salary: 60000}

	store.Create(emp)

	updated := store.Update(emp.ID, newEmp)
	assert.True(t, updated)

	result, _ := store.GetByID(emp.ID)
	assert.Equal(t, newEmp, result)
}

func TestDeleteEmployee(t *testing.T) {
	store := NewInMemoryStore()
	emp := employee.Employee{ID: 1, Name: "Chetan Tiwari", Position: "Software Engineer", Salary: 50000}

	store.Create(emp)

	deleted := store.Delete(emp.ID)
	assert.True(t, deleted)

	_, ok := store.GetByID(emp.ID)
	assert.False(t, ok)
}

func TestListEmployee(t *testing.T) {
	store := NewInMemoryStore()
	emp1 := employee.Employee{ID: 1, Name: "Chetan Tiwari", Position: "Software Engineer", Salary: 50000}
	emp2 := employee.Employee{ID: 2, Name: "Prince Sharma", Position: "Product Manager", Salary: 60000}
	emp3 := employee.Employee{ID: 3, Name: "Vijay Pandey", Position: "Data Scientist", Salary: 70000}

	store.Create(emp1)
	store.Create(emp2)
	store.Create(emp3)

	page := 1
	pageSize := 1
	employees := store.List(page, pageSize)

	assert.Len(t, employees, 1)

	expectedEmployees := []employee.Employee{emp1}
	assert.Equal(t, expectedEmployees, employees)
}

// edge cases

func TestCreateDuplicateEmployee(t *testing.T) {
	// Test creating an employee with a duplicate ID
	store := NewInMemoryStore()
	emp := employee.Employee{ID: 1, Name: "Chetan Tiwari", Position: "Software Engineer", Salary: 50000}

	// Add employee to store
	store.Create(emp)

	// Try adding the same employee again
	store.Create(emp)

	// Employee should still exist in the store
	_, ok := store.GetByID(emp.ID)
	assert.True(t, ok)
}

func TestUpdateNonExistentEmployee(t *testing.T) {
	// Test updating details of an employee that doesn't exist
	store := NewInMemoryStore()
	newEmp := employee.Employee{ID: 1, Name: "Chetan Tiwari", Position: "Senior Engineer", Salary: 60000}

	// Try updating details of non-existent employee
	updated := store.Update(1, newEmp)
	assert.False(t, updated)
}

func TestDeleteNonExistentEmployee(t *testing.T) {
	// Test deleting an employee that doesn't exist
	store := NewInMemoryStore()

	// Try deleting non-existent employee
	deleted := store.Delete(100)
	assert.False(t, deleted)
}

func TestListWithEmptyStore(t *testing.T) {
	// Test listing employees when the store is empty
	store := NewInMemoryStore()

	// List employees with pagination
	page := 1
	pageSize := 10
	employees := store.List(page, pageSize)

	// There should be no employees in the list
	assert.Empty(t, employees)
}

func TestListWithInvalidPage(t *testing.T) {
	// Test listing employees with an invalid page number
	store := NewInMemoryStore()

	// Add some employees to the store
	emp1 := employee.Employee{ID: 1, Name: "Chetan Tiwari", Position: "Software Engineer", Salary: 50000}
	emp2 := employee.Employee{ID: 2, Name: "Prince Sharma", Position: "Product Manager", Salary: 60000}
	emp3 := employee.Employee{ID: 3, Name: "Vijay Pandey", Position: "Data Scientist", Salary: 70000}

	store.Create(emp1)
	store.Create(emp2)
	store.Create(emp3)

	// List employees with an invalid page number
	page := -1
	pageSize := 2
	employees := store.List(page, pageSize)

	// There should be no employees in the list
	assert.Empty(t, employees)
}

func TestListWithLargeDataset(t *testing.T) {
	// Test listing employees with a large dataset
	store := NewInMemoryStore()
	const employeesCnt = 1000

	for i := 1; i <= employeesCnt; i++ {
		emp := employee.Employee{ID: i, Name: fmt.Sprintf("Employee %d", i), Position: "Position", Salary: float64(i * 1000)}
		store.Create(emp)
	}

	// List employees with pagination
	page := 1
	pageSize := 100
	employees := store.List(page, pageSize)

	assert.Len(t, employees, pageSize)
}
