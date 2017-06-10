// From https://play.golang.org/p/9U22NfrXeq

package primes

// Send the sequence 2, 3, 5, 7, 9 ... to channel 'ch'.
func generate(ch chan<- int) {
    ch <- 2
	for i := 3; ; i+=2 {
		ch <- i
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i % prime != 0 {
			out<- i
		}
	}
}

func Channel() (chan int) {
    out := make(chan int)
    ch := make(chan int)
    go generate(ch)

    go func() {
        for {
            prime := <-ch
            out <-prime
            ch1 := make(chan int)
            go filter(ch, ch1, prime)
            ch = ch1
        }
    }()
    
    return out
}
