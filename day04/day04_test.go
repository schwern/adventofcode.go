package day04_test

import(
    "testing"
    "github.com/schwern/adventofcode2015/testutil"
    "github.com/schwern/adventofcode2015/day04"
)

var Input_File = "../testdata/day04.txt"

func TestMineAdventCoin( t *testing.T ) {
    tests := []testutil.TestCase{
        { "abcdef", 609043 },
        { "pqrstuv", 1048970 },
        { "ckczppom", 117946 },
    }
    
    for _, test := range tests {
        have := day04.MineAdventCoin( test.Arg )
        testutil.AssertEq( t, have, test.Want )
    }
}
