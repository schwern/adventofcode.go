package day18_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/2015/day18"
)

func TestStep( t *testing.T ) {
    lines := [][]string{
        {
            ".#.#.#",
            "...##.",
            "#....#",
            "..#...",
            "#.#..#",
            "####..",
        },
        {
            "..##..",
            "..##.#",
            "...##.",
            "......",
            "#.....",
            "#.##..",
        },
        {
            "..###.",
            "......",
            "..###.",
            "......",
            ".#....",
            ".#....",
        },
        {
            "...#..",
            "......",
            "...#..",
            "..##..",
            "......",
            "......",
        },
        {
            "......",
            "......",
            "..##..",
            "..##..",
            "......",
            "......",
        },
    }
    
    states := make( []day18.Grid, 0 )
    for _,line := range lines {
        grid := day18.ParseGrid( line )
        states = append(states, grid)
    }
    
    gol := day18.NewGOL( 6, 6, states[0] )
    for _,state := range states {
        testutil.AssertEq( t, gol.Grid(), state )
        gol.Step()
    }
    
    // The final state is stable, so it's ok we went an extra step.
    testutil.AssertEq( t, gol.HowManyLightsDoYouSee(), 4 )
}
