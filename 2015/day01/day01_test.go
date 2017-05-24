package day01_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/2015/day01"
    "github.com/schwern/adventofcode.go/util"
)

var input_file = "testdata/input.txt"

func TestFindFloor(t *testing.T) {
    tests := []testutil.TestCase {
        {"(())", 0},
        {"()()", 0},
        {"(((", 3},
        {"(()(()(", 3},
        {"))(((((", 3},
        {"())", -1},
        {"))(", -1},
        {")))", -3},
        {")())())", -3},
    }
    
    input := util.ReadFile(input_file)
    tests = append( tests, testutil.TestCase{ input, 138 } )
    
    for _, test := range tests {
        have := day01.FindFloor(test.Arg)
        testutil.AssertEq( t, have, test.Want )
    }
}

func TestFirstBasement(t *testing.T) {
    tests := []testutil.TestCase {
        { ")", 1 },
        { "()())", 5 },
        { "((", 0 },
    }
    
    input := util.ReadFile(input_file)
    tests = append( tests, testutil.TestCase{ input, 1771 } )
    
    for _, test := range tests {
        have := day01.FirstBasement(test.Arg)
        testutil.AssertEq( t, have, test.Want )
    }
}
