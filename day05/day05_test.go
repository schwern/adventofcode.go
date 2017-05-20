package day05_test

import(
    "testing"
    //"github.com/schwern/adventofcode2015/util"
    "github.com/schwern/adventofcode2015/testutil"
    "github.com/schwern/adventofcode2015/day05"
)

var Input_File = "../testdata/day05.txt"

func TestIsNice( t *testing.T ) {
    tests := []struct{ Arg string; Want bool }{
        { "ugknbfddgicrmopn", true },
        { "aaa", true },
        { "jchzalrnumimnmhp", false },
        { "haegwjzuvuyypxyu", false },
        { "dvszwmarrgswjxmb", false },
    }
    
    for _, test := range tests {
        have := day05.IsNice(test.Arg)
        testutil.AssertEq( t, have, test.Want )
    }
}

func TestIsNice2( t *testing.T ) {
    tests := []struct{ Arg string; Want bool }{
        { "qjhvhtzxzqqjkmpb", true },
        { "xxyxx", true },
        { "uurcxstgmygtbstg", false },
        { "ieodomkazucvgmuy", false },
    }
    
    for _, test := range tests {
        have := day05.IsNice2(test.Arg)
        testutil.AssertEq( t, have, test.Want )
    }
}
