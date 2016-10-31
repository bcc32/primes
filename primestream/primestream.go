package primestream

func filter(dst chan<- int, src <-chan int, modulo int) {
    for n := range src {
        if n % modulo != 0 {
            dst <- n
        }
    }
}

func nats(dst chan<- int, start int) {
    for i := start; ; i++ {
        dst <- i
    }
}

func Primes(dst chan<- int) {
    stream := make(chan int, 1000)
    go nats(stream, 2)
    for {
        next := <-stream
        dst <- next
        filtered := make(chan int, 100)
        go filter(filtered, stream, next)
        stream = filtered
    }
}
