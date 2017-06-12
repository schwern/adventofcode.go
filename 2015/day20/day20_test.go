package day20_test

import(
    "testing"
    "github.com/stvp/assert"
    "github.com/schwern/adventofcode.go/2015/day20"
)

func TestNumPresents( t *testing.T ) {
    want := []int{
        0, 10, 30, 40, 70, 60, 120, 80, 150, 130,
    }

    assert.Equal( t, day20.DeliverPresents(9), want )
}

func TestPart1( t *testing.T ) {
    want := 29000000
    
    // Worst case is a prime number gets (n+1)*10 presents.
    // So want/10 is the highest we need to go.
    presents := day20.DeliverPresents(want/10)
    for house, num := range presents {
        if num > want {
            assert.Equal( t, house, 665280 )
            break
        }
    }
}
