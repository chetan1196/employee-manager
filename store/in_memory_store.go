package store

import (
	"employee-manager/employee"
	"log"
	"sync"
)

// inMemoryStore implements the Store interface
type inMemoryStore struct {
	sync.RWMutex
	employees map[int]employee.Employee
	idCounter int // idCounter for generating unique id
}

// NewInMemoryStore creates a new instance of inMemoryStore
func NewInMemoryStore() EmployeeStore {
	return &inMemoryStore{
		employees: make(map[int]employee.Employee),
		idCounter: 1,
	}
}

// GenerateID generates a unique ID for new employees
func (ims *inMemoryStore) GenerateID() int {
	ims.Lock()
	defer ims.Unlock()

	id := ims.idCounter
	ims.idCounter++
	return id
}

// Create creates new employees
func (ims *inMemoryStore) Create(emp employee.Employee) {
	ims.Lock()
	defer ims.Unlock()
	ims.employees[emp.ID] = emp
}

// GetByID returns employee by its id
func (ims *inMemoryStore) GetByID(id int) (employee.Employee, bool) {
	ims.RLock()
	defer ims.RUnlock()
	emp, ok := ims.employees[id]
	return emp, ok
}

// Update updates the employee using id
func (ims *inMemoryStore) Update(id int, newEmp employee.Employee) bool {
	ims.Lock()
	defer ims.Unlock()
	if _, ok := ims.employees[id]; ok {
		ims.employees[id] = newEmp
		return true
	}
	return false
}

// Delete delete the employee from the store
func (ims *inMemoryStore) Delete(id int) bool {
	ims.Lock()
	defer ims.Unlock()
	if _, ok := ims.employees[id]; ok {
		delete(ims.employees, id)
		return true
	}
	return false
}

// List listed the existing employees. It supports pagination.
// If no pagination paramerters or invalid parameters are provided, List default to default parameters
func (ims *inMemoryStore) List(page, pageSize int) []employee.Employee {
	ims.RLock()
	defer ims.RUnlock()

	// Calculate the starting and ending index for pagination
	start := (page - 1) * pageSize
	end := start + pageSize
	if start >= len(ims.employees) {
		return []employee.Employee{}
	}

	if end > len(ims.employees) {
		end = len(ims.employees)
	}

	log.Printf("start: %v, end: %v", start, end)

	// Extract the employees for the current page
	employees := make([]employee.Employee, 0, pageSize)
	i := 0
	for _, emp := range ims.employees {
		if i > end {
			break
		}
		if i >= start && i < end {
			employees = append(employees, emp)
		}
		i++
	}
	return employees
}
