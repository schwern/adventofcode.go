package primes

import(
    "math"
)

var primes = []int{2,3,5,7,11,13,17}

// Returns a slice of prime numbers at least up to `upto`.
// The returned slice is shared, don't modify it.
func UpTo( upto int ) []int {
    if upto <= primes[len(primes)-1] {
        return primes
    }
    
    // Grow the array by what should be a sufficient amount.
    // Using x/log(x) as the number of primes < x.
    approxPrimes := float64(upto) / math.Log(float64(upto))
    newCap := int(approxPrimes) + 1
    if newCap > cap(primes) {
        tmp := make([]int, len(primes), newCap)
        copy(tmp, primes)
        primes = tmp
    }
    
    for c := primes[len(primes)-1] + 2; upto > c; c+= 2 {
        isprime := true
        for _,prime := range primes {
            if prime*2 > c {
                break
            }
            if c % prime == 0 {
                // it's composite
                isprime = false
                break
            }
        }
        
        if isprime {
            primes = append(primes, c)
        }
    }
    
    return primes
}
