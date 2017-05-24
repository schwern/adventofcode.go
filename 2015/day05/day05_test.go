package day05_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/util"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/2015/day05"
)

var Input_File = "testdata/input.txt"

type checkFunc func( string ) bool

func checkList( checker checkFunc ) int {
    lines := util.LineChannel(Input_File)
    
    num_nice := 0
    for line := range lines {
        if checker(line) {
            num_nice++
        }
    }
    
    return num_nice
}

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
    
    testutil.AssertEq( t, checkList( day05.IsNice ), 236 )
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
    
    testutil.AssertEq( t, checkList( day05.IsNice2 ), 51 )
}
