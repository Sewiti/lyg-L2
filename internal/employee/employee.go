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
	const format = "%-30s|%4d |%9.2f"

	if e.Hash == "" {
		return fmt.Sprintf(format, e.Name, e.Age, e.Salary)
	}

	return fmt.Sprintf(format+" | %s", e.Name, e.Age, e.Salary, e.Hash)
}
