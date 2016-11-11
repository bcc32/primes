package primestream

func filter(src <-chan int, modulo int) <-chan int {
	dst := make(chan int)
	go func() {
		for n := range src {
			if n%modulo != 0 {
				dst <- n
			}
		}
	}()
	return dst
}

func nats(start int) <-chan int {
	dst := make(chan int)
	go func() {
		for i := start; ; i++ {
			dst <- i
		}
	}()
	return dst
}

func Primes() <-chan int {
	dst := make(chan int)
	stream := nats(2)
	go func() {
		for {
			next := <-stream
			dst <- next
			filtered := filter(stream, next)
			stream = filtered
		}
	}()
	return dst
}
