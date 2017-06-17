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
    assert.Equal( t, combos, [][]int{ {7,9} } )
    
    combos = tryCombos( nums, 3, 20 )
    assert.Equal(
        t, combos,
        [][]int{
            {0,6,9}, {0,7,8},
            {1,5,9}, {1,6,8},
            {2,5,8}, {2,6,7},
            {3,4,9}, {3,5,7},
            {4,5,6},
        },
    )
}

func TestSmallestBucketsChan( t *testing.T ) {
    nums := []int{1,2,3,4,5,7,8,9,10,11}

    wants := [][]int{
        {7,9},
        {0,6,9}, {0,7,8},
        {1,5,9}, {1,6,8},
        {2,5,8}, {2,6,7},
        {3,4,9}, {3,5,7},
        {4,5,6},
    }
    
    bucketChan := smallestBucketsChan( nums )
    for _,want := range wants {
        assert.Equal( t, <-bucketChan, want )
    }
}
