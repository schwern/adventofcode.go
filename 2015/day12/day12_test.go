package day12_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/2015/day12"
)

func TestSumNums( t *testing.T ) {
    tests := []testutil.TestCase{
        { "[1,2,3]", 6 },
        { `{"a":2,"b":4}`, 6 },
        { `[[[3]]]`, 3 },
        { `{"a":{"b":4},"c":-1}`, 3 },
        { `{"a":[-1,1]}`, 0 },
        { `[-1,{"a":1}]`, 0 },
        { `[]`, 0 },
        { `{}`, 0 },
    }
    
    for _,test := range tests {
        have := day12.SumNums( test.Arg )
        testutil.AssertEq( t, have, test.Want )
    }
}
