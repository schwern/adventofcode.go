package day09_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/2015/day09"
)

func TestPermutation( t *testing.T ) {
    tests := []struct{ arg []int; want [][]int }{
        { []int{}, [][]int{} },
        { []int{11}, [][]int{ {11} } },
        { []int{11, 22}, [][]int{ {11,22}, {22,11} } },
        {
            []int{11, 22, 33},
            [][]int{ {11,22,33}, {11,33,22},
                     {22,11,33}, {22,33,11},
                     {33,22,11}, {33,11,22},
            },
        },
    }

    for _, test := range tests {
        perm := day09.NewPermutation( test.arg )
    
        i := 0
        for next := perm.Next(); next != nil; next = perm.Next() {
            testutil.AssertIntSliceEq( t, next, test.want[i] )
            i++
        }

        i = 0
        ch := day09.NewPermutationChan( test.arg )
        for set := range ch {
            testutil.AssertIntSliceEq( t, set, test.want[i] )
            i++
        }
    }
}
