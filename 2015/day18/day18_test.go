package day18_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/util"
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

func TestPart1( t *testing.T ) {
    grid := day18.ParseGridChan( util.LineChannel( InputFile ) )
    gol := day18.NewGOL( 100, 100, grid )
    
    for i := 0; i < 100; i++ {
        gol.Step()
    }
    
    testutil.AssertEq( t, gol.HowManyLightsDoYouSee(), 821 )
}
