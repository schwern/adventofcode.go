package day03_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/2015/day03"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/util"
)

var Input_File = "testdata/input.txt"

func TestPresentsDelivered( t *testing.T ) {
    tests := []testutil.TestCase{
        { ">", 2 },
        { "^>v<", 4 },
        { "^v^v^v^v^v", 2 },
        { util.ReadFile(Input_File), 2592 },
    }

    for _, test := range tests {
        have := day03.DeliverPresents( test.Arg, 1 )
        testutil.AssertEq( t, have, test.Want )
    }
}

func TestRoboPresentsDelivered( t *testing.T ) {
    tests := []testutil.TestCase{
        { "^v", 3 },
        { "^>v<", 3 },
        { "^v^v^v^v^v", 11 },
        { util.ReadFile(Input_File), 2360 },
    }

    for _, test := range tests {
        have := day03.DeliverPresents( test.Arg, 2 )
        testutil.AssertEq( t, have, test.Want )
    }
}
