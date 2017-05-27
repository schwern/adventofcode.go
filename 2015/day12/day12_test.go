package day12_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/util"
    "github.com/schwern/adventofcode.go/2015/day12"
)

var Input_File = "testdata/input.txt"

func TestSumNums( t *testing.T ) {
    tests := []struct{ Arg string; Nil float64; Red float64 }{
        { "[1,2,3]", 6, 6 },
        { `{"a":2,"b":4}`, 6, 6 },
        { `[[[3]]]`, 3, 3 },
        { `{"a":{"b":4},"c":-1}`, 3, 3 },
        { `{"a":[-1,1]}`, 0, 0 },
        { `[-1,{"a":1}]`, 0, 0 },
        { `[]`, 0, 0 },
        { `{}`, 0, 0 },
        { `[1,{"c":"red","b":2},3]`, 6, 4 },
        { `{"d":"red","e":[1,2,3,4],"f":5}`, 15, 0 },
        { `[1,"red",5]`, 6, 6 },
    }
    
    for _,test := range tests {
        have := day12.SumNums( test.Arg, nil )
        testutil.AssertEq( t, have, test.Nil )

        skip := "red"
        have = day12.SumNums( test.Arg, &skip )
        testutil.AssertEq( t, have, test.Red )
    }
}

func TestPart1( t *testing.T ) {
    have := day12.SumNums( util.ReadFile( Input_File ), nil )
    testutil.AssertEq( t, have, float64(119433) )
}

func TestPart2( t *testing.T ) {
    skip := "red"
    have := day12.SumNums( util.ReadFile( Input_File ), &skip )
    testutil.AssertEq( t, have, float64(68466) )
}
