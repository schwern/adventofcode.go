package combination_test

import(
    "testing"
    "github.com/stvp/assert"
    "github.com/schwern/adventofcode.go/combination"
)

func TestCombination( t *testing.T ) {
    want := [][]int{
        {0,1,2},
        {0,1,3},
        {0,2,3},
        {1,2,3},
    }
    
    i := 0
    for combos := range combination.Chan(4,3) {
        assert.Equal( t, combos, want[i] )
        i++
    }
}
