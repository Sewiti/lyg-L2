package employee

import "fmt"

// Employee holds data about employee such as name, age, salary & computed hash
type Employee struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Salary float64 `json:"salary"`
	Hash   string
}

func (e *Employee) String() string {
	return fmt.Sprintf("%s %d %f %s", e.Name, e.Age, e.Salary, e.Hash)
}
