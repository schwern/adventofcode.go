package day04_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/2015/day04"
)

var Input_File = "testdata/input.txt"

func TestMineAdventCoin( t *testing.T ) {
    tests := []testutil.TestCase{
        { "abcdef", 609043 },
        { "pqrstuv", 1048970 },
        // Part 1
        { "ckczppom", 117946 },
    }
    
    for _, test := range tests {
        have := day04.MineAdventCoin( test.Arg, "00000" )
        testutil.AssertEq( t, have, test.Want )
    }
    
    // Part 2
    testutil.AssertEq(
        t,
        day04.MineAdventCoin( "ckczppom", "000000" ),
        3938038,
    )
}
