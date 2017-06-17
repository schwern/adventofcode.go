package day24

import(
    "testing"
    "github.com/stvp/assert"
)

func TestTryCombos( t *testing.T ) {
    nums := []int{1,2,3,4,5,7,8,9,10,11}
    
    combos := tryCombos( nums, 1, 20 )
    assert.Equal( t, combos, [][]int{} )
    
    combos = tryCombos( nums, 2, 20 )
    assert.Equal( t, combos, [][]int{ {9,11} } )
    
    combos = tryCombos( nums, 3, 20 )
    assert.Equal(
        t, combos,
        [][]int{
            {1,8,11}, {1,9,10},
            {2,7,11}, {2,8,10},
            {3,7,10}, {3,8,9},
            {4,5,11}, {4,7,9},
            {5,7,8},
        },
    )
}
