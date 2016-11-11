package main

import (
	"fmt"
	"github.com/bcc32/primes/primestream"
)

func main() {
	for n := range primestream.Primes() {
		fmt.Println(n)
	}
}
