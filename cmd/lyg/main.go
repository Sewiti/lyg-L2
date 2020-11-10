package main

import (
	"fmt"
	"math"

	"github.com/Sewiti/lyg-L2/internal/employee"
)

func main() {
	fileSet := []struct {
		in  string
		out string
	}{
		{"data/data-1.json", "output/res-1.txt"},
		{"data/data-2.json", "output/res-2.txt"},
		{"data/data-3.json", "output/res-3.txt"},
	}

	for _, fs := range fileSet {
		fmt.Printf("Executing '%s'...\n", fs.in)

		if err := execute(fs.in, fs.out); err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Done! Have a nice day:)")
}

func execute(dataFile string, resFile string) error {
	es, err := read(dataFile)
	if err != nil {
		return err
	}

	n := int(math.Max(2, float64(len(es))/4)) // workers count
	m := int(math.Max(1, float64(len(es))/2)) // data thread internal array size

	snd := make(chan employee.Employee)    // for sending to data thread
	get := make(chan employee.Employee, n) // for retrieving from data thread
	res := make(chan employee.Employee, n) // for sending to results thread
	fin := make(chan []employee.Employee)  // for retrieving from results thread
	sync := make(chan struct{})            // for synchronization

	for i := 0; i < n; i++ {
		go worker(get, res, sync)
	}

	go data(snd, get, m)
	go result(res, fin)

	for _, e := range es {
		snd <- e
	}
	close(snd)

	for i := 0; i < n; i++ {
		<-sync
	}
	close(sync)
	close(res)

	return write(resFile, es, <-fin)
}
