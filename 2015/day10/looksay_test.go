package day10_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/2015/day10"
)

var Input = "1113122113"

func TestLookSay( t *testing.T ) {
    tests := []struct{ arg []byte; want []byte }{
        { []byte("1"), []byte("11") },
        { []byte("11"), []byte("21") },
        { []byte("21"), []byte("1211") },
        { []byte("1211"), []byte("111221") },
        { []byte("111221"), []byte("312211") },
    }
    
    for _, test := range tests {
        testutil.AssertEq(
            t,
            string( day10.LookSay( test.arg ) ),
            string( test.want),
        )
    }
}

func TestPart1And2( t *testing.T ) {
    // Part 1
    have := []byte(Input)
    for i := 0; i < 40; i++ {
        have = day10.LookSay( have )
    }
    testutil.AssertEq( t, len(have), 360154 )
    
    // Part 2
    for i := 0; i < 10; i++ {
        have = day10.LookSay( have )
    }
    testutil.AssertEq( t, len(have), 5103798 )
}
