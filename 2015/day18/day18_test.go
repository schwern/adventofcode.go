package day18_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/2015/day18"
)

var InputFile = "testdata/input.txt"

func TestParseGrid( t *testing.T ) {
    lines := []string{
        ".#.#.#",
        "...##.",
        "#....#",
        "..#...",
        "#.#..#",
        "####..",
    }
    
    want := day18.Grid{
        {false, true, false, true, false, true},
        {false, false, false, true, true, false},
        {true, false, false, false, false, true},
        {false, false, true, false, false, false},
        {true, false, true, false, false, true},
        {true, true, true, true, false, false},
    }
    
    testutil.AssertEq( t, day18.ParseGrid(lines), want )
}
