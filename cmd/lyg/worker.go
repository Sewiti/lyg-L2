package main

import (
	"crypto/sha512"
	"encoding/base64"
	"sync"

	"github.com/Sewiti/lyg-L2/internal/employee"
)

const iterations = 6e6

func worker(in <-chan employee.Employee, out chan<- employee.Employee, wg *sync.WaitGroup) {
	defer wg.Done()

	for e := range in {
		bytes := []byte(e.String())

		hasher := sha512.New()
		for i := 0; i < iterations; i++ {
			hasher.Write(bytes)
		}

		e.Hash = base64.URLEncoding.EncodeToString(hasher.Sum(nil))

		if e.Age >= 18 {
			out <- e
		}
	}
}
