package main

import (
	"reflect"

	"github.com/Sewiti/lyg-L2/internal/employee"
)

func data(in <-chan employee.Employee, out chan<- employee.Employee, max int) {
	arr := make([]employee.Employee, max)
	from, to, size := 0, 0, 0
	done := false

	for {
		var cases []reflect.SelectCase

		if size < len(arr) {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(in),
			})
		} else {
			// Placeholder to keep indexes in check
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Send: reflect.Value{},
			})
		}

		if size > 0 {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectSend,
				Chan: reflect.ValueOf(out),
				Send: reflect.ValueOf(arr[from]),
			})
		} else if done {
			break
		}

		switch i, v, ok := reflect.Select(cases); i {
		case 0:
			if ok {
				arr[to] = v.Interface().(employee.Employee)
				to = (to + 1) % len(arr)
				size++
			} else if !done {
				done = true
			}

		case 1:
			arr[from] = employee.Employee{} // optional
			from = (from + 1) % len(arr)
			size--
		}
	}
}
