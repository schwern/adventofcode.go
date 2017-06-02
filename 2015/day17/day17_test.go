package day17_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/util"
    "github.com/schwern/adventofcode.go/2015/day17"
)

var InputFile = "testdata/input.txt"

func TestNumCombos( t *testing.T ) {
    containers := day17.Containers{20, 15, 10, 5, 5}
    testutil.AssertEq( t, containers.NumCombos(25), 4 )
}

func TestPart1( t *testing.T ) {
    lines := util.LineChannel(InputFile)
    containers := day17.Containers{}
    for line := range lines {
        containers = append( containers, util.MustAtoi(line) )
    }
    
    testutil.AssertEq( t, containers.NumCombos(150), 1638 )
}
