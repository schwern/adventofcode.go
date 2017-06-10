package primes_test

import(
    "testing"
    "github.com/stvp/assert"
    "github.com/schwern/adventofcode.go/primes"
)

func TestChannel( t *testing.T ) {
    wants := []int{ 2, 3, 5, 7, 11, 13, 17, 19, 23, 29 }
    pChan := primes.Channel()
    
    for _,want := range wants {
        assert.Equal( t, <-pChan, want )
    }
}
