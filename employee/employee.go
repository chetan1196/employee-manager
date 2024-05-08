package employee

// Employee represents an employee entity
type Employee struct {
	ID       int     `json:"-"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}
