package day10_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/2015/day10"
)

var Input = "1113122113"

func TestLookSay( t *testing.T ) {
    tests := []struct{ arg string; want string }{
        { "1", "11" },
        { "11", "21" },
        { "21", "1211" },
        { "1211", "111221" },
        { "111221", "312211" },
    }
    
    for _, test := range tests {
        testutil.AssertEq( t, day10.LookSay( test.arg ), test.want )
    }
}

func TestPart1( t *testing.T ) {
    have := Input
    for i := 0; i < 40; i++ {
        have = day10.LookSay( have )
    }
    
    testutil.AssertEq( t, len(have), 360154 )
}