package day24_test

import(
    "testing"
    "github.com/stvp/assert"
    "github.com/schwern/adventofcode.go/2015/day24"
)

var Packs = []int{
    1, 2, 3, 5, 7, 13, 17, 19, 23, 29, 31, 37, 41, 43, 53,
    59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
}

func TestFindSmallestCombos( t *testing.T ) {
    nums := []int{1,2,3,4,5,7,8,9,10,11}

    assert.Equal(
        t, day24.FindSmallestCombos( nums ), [][]int{ {9,11} },
    )
}

func TestSmallestQE( t *testing.T ) {
    assert.Equal(
        t, day24.SmallestQE( [][]int{ {10,9,1}, {10,8,2} } ), 90,
    )
}
