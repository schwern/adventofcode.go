package primes_test

import(
    "testing"
    "github.com/stvp/assert"
    "github.com/schwern/adventofcode.go/primes"
)

func TestUpTo( t *testing.T ) {
    want := []int{ 2, 3, 5, 7, 11, 13, 17, 19, 23, 29 }
    
    have := primes.UpTo( 30 )    
    assert.Equal( t, have, want )
    
    have = primes.UpTo( 0 )
    assert.Equal( t, have, want )
}
