package store

import (
	"employee-manager/employee"
)

// EmployeeStore defines the interface for managing employees
type EmployeeStore interface {
	GenerateID() int
	Create(emp employee.Employee)
	GetByID(id int) (employee.Employee, bool)
	Update(id int, newEmp employee.Employee) bool
	Delete(id int) bool
	List(page, pageSize int) []employee.Employee
}
