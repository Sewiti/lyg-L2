package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Sewiti/lyg-L2/internal/employee"
)

func read(fn string) ([]employee.Employee, error) {
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var e []employee.Employee
	err = json.NewDecoder(file).Decode(&e)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func write(fn string, init []employee.Employee, res []employee.Employee) error {
	file, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer file.Close()

	if err = writeSection(file, "Initial data:", init, false); err != nil {
		return err
	}

	return writeSection(file, "\nResults:", res, true)
}

func writeSection(file *os.File, title string, es []employee.Employee, full bool) error {
	const (
		header     = "  # | Name                          | Age | Salary"
		headerFull = header + "   | Hash"

		empty     = "  - | -                             |   - |    --.--"
		emptyFull = empty + " | -"

		line     = "----+-------------------------------+-----+----------"
		lineFull = line + "+------------------------------------------------------------------------------------------"
	)

	var hl, em string
	if full {
		hl = headerFull + "\n" + lineFull + "\n"
		em = emptyFull + "\n"
	} else {
		hl = header + "\n" + line + "\n"
		em = empty + "\n"
	}

	if _, err := file.WriteString(fmt.Sprintf("%s\n%s", title, hl)); err != nil {
		return err
	}

	if len(es) > 0 {
		for i, e := range es {
			if _, err := file.WriteString(fmt.Sprintf("%3d | %s\n", i+1, e.String())); err != nil {
				return err
			}
		}
	} else {
		if _, err := file.WriteString(em); err != nil {
			return err
		}
	}

	return nil
}
