package employee

// Employee holds data about employee such as name, age, salary & computed hash
type Employee struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Salary float64 `json:"salary"`
	Hash   string
}
