package day03_test

import(
    "testing"
    "github.com/schwern/adventofcode2015/day03"
    "github.com/schwern/adventofcode2015/testutil"
    "github.com/schwern/adventofcode2015/util"
)

var Input_File = "../testdata/day03.txt"

func TestPresentsDelivered( t *testing.T ) {
    tests := []testutil.TestCase{
        { ">", 2 },
        { "^>v<", 4 },
        { "^v^v^v^v^v", 2 },
        { util.ReadFile(Input_File), 2592 },
    }

    for _, test := range tests {
        have := day03.PresentsDelivered( test.Arg )
        testutil.AssertEq( t, have, test.Want )
    }
}
