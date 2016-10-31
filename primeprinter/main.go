package main

import (
    "fmt"
    "github.com/bcc32/primes/primestream"
)

func main() {
    primes := make(chan int, 100)
    go primestream.Primes(primes)
    for n := range primes {
        fmt.Println(n)
    }
}
