package day14_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/2015/day14"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/util"
)

var Input_File = "testdata/input.txt"

var Comet  = day14.Reindeer{ Speed: 14, Duration: 10, Rest: 127 }
var Dancer = day14.Reindeer{ Speed: 16, Duration: 11, Rest: 162 }

func TestRunRunReindeer( t *testing.T ) {
    tests := []struct{ Arg day14.Reindeer; Want int }{
        { Comet, 1120 },
        { Dancer, 1056 },
    }
    
    for _,test := range tests {
        have := test.Arg.RunRunReindeer( 1000 )
        testutil.AssertEq( t, have, test.Want )
    }
}

func TestParseLine( t *testing.T ) {
    tests := []struct{ Arg string; Want day14.Reindeer }{
        {
            Arg: `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.`,
            Want: Comet,
        },
        {
            Arg: `Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`,
            Want: Dancer,
        },
    }
    
    for _,test := range tests {
        have := day14.ParseLine( test.Arg )
        testutil.AssertEq( t, have, test.Want )
    }
}

func TestPart1( t *testing.T ) {
    lines := util.LineChannel( Input_File )
    
    time := 2503
    
    dist := 0
    for line := range lines {
        r := day14.ParseLine( line )
        dist = util.MaxInt( r.RunRunReindeer( time ), dist )
    }
    
    testutil.AssertEq( t, dist, 2640 )
}

func TestReindeerRace( t *testing.T ) {
    have := day14.ReindeerRace(
        []day14.Reindeer{ Comet, Dancer },
        140,
    )
    want := []int{1,139}
    
    testutil.AssertIntSliceEq( t, have, want )
    
    have = day14.ReindeerRace(
        []day14.Reindeer{ Comet, Dancer },
        1000,
    )
    want = []int{312,689}
    
    testutil.AssertIntSliceEq( t, have, want )
}

func TestPart2( t *testing.T ) {
    lines := util.LineChannel( Input_File )
    
    time := 2503

    rs := make([]day14.Reindeer, 0)
    for line := range lines {
        r := day14.ParseLine( line )
        rs = append( rs, r )
    }
    points := day14.ReindeerRace( rs, time )
    
    testutil.AssertEq( t, util.MaxInts(points), 1102 )
}
