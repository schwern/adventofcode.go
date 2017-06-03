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
    for i := 1; i < len(states); i++ {
        gol.Step()
        testutil.AssertEq( t, gol.Grid(), states[i] )
    }
    
    testutil.AssertEq( t, gol.HowManyLightsDoYouSee(), 4 )
}

func TestStuck( t *testing.T ) {
    lines := [][]string{
        {
            "##.#.#",
            "...##.",
            "#....#",
            "..#...",
            "#.#..#",
            "####.#",
        },
        {
            "#.##.#",
            "####.#",
            "...##.",
            "......",
            "#...#.",
            "#.####",
        },
        {
            "#..#.#",
            "#....#",
            ".#.##.",
            "...##.",
            ".#..##",
            "##.###",
        },
        {
            "#...##",
            "####.#",
            "..##.#",
            "......",
            "##....",
            "####.#",
        },
        {
            "#.####",
            "#....#",
            "...#..",
            ".##...",
            "#.....",
            "#.#..#",
        },
        {
            "##.###",
            ".##..#",
            ".##...",
            ".##...",
            "#.#...",
            "##...#",
        },
    }
    
    stuck := []string{
        "#....#",
        "......",
        "......",
        "......",
        "......",
        "#....#",
    }
    
    states := make( []day18.Grid, 0 )
    for _,line := range lines {
        grid := day18.ParseGrid( line )
        states = append(states, grid)
    }
    
    gol := day18.NewGOL( 6, 6, states[0] )
    gol.AddStuckLights( day18.ParseGrid(stuck) )
    for i := 1; i < len(states); i++ {
        gol.Step()
        testutil.AssertEq( t, gol.Grid(), states[i] )
    }
    
    // The final state is stable, so it's ok we went an extra step.
    testutil.AssertEq( t, gol.HowManyLightsDoYouSee(), 17 )
}
