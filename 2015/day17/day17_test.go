package day17_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/2015/day17"
)

func TestNumCombos( t *testing.T ) {
    containers := day17.Containers{20, 15, 10, 5, 5}
    testutil.AssertEq( t, containers.NumCombos(25), 4 )
}
