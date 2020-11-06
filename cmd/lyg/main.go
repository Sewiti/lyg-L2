package main

import (
	"encoding/json"
	"os"

	"github.com/Sewiti/lyg-L2/internal/employee"
)

func main() {
	execute("data/data-1.json")
}

func execute(filename string) error {
	employees, err := read(filename)
	if err != nil {
		return err
	}

	send := make(chan employee.Employee)
	get := make(chan employee.Employee)

	go data(send, get, len(employees))

	for _, e := range employees {
		send <- e
	}
	close(send)

	return nil
}

func read(filename string) ([]employee.Employee, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var e []employee.Employee
	err = json.NewDecoder(file).Decode(&e)
	if err != nil {
		return nil, err
	}

	return e, err
}
