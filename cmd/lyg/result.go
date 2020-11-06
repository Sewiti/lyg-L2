package main

import "github.com/Sewiti/lyg-L2/internal/employee"

func result(in <-chan employee.Employee, out chan<- []employee.Employee) {
	var arr []employee.Employee

	for e := range in {
		var i int
		for i = 0; i < len(arr); i++ {
			if arr[i].Name >= e.Name {
				break
			}
		}

		arr = append(arr, employee.Employee{})
		copy(arr[i+1:], arr[i:])
		arr[i] = e
	}

	out <- arr
	close(out)
}
