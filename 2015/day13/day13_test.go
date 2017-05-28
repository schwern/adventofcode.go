package day13_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/2015/day13"
    "github.com/schwern/adventofcode.go/routes"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/util"
)

var Input_File = "testdata/input.txt"

var Tests = []struct{ Arg string; a,b string; dist int }{
    {
        "Alice would gain 54 happiness units by sitting next to Bob.",
        "Alice", "Bob", 54,
    },
    {
        "Alice would lose 79 happiness units by sitting next to Carol.",
        "Alice", "Carol", -79,
    },
    {
        "Alice would lose 2 happiness units by sitting next to David.",
        "Alice", "David", -2,
    },
    {
        "Bob would gain 83 happiness units by sitting next to Alice.",
        "Bob", "Alice", 83,
    },
    {
        "Bob would lose 7 happiness units by sitting next to Carol.",
        "Bob", "Carol", -7,
    },
    {
        "Bob would lose 63 happiness units by sitting next to David.",
        "Bob", "David", -63,
    },
    {
        "Carol would lose 62 happiness units by sitting next to Alice.",
        "Carol", "Alice", -62,
    },
    {
        "Carol would gain 60 happiness units by sitting next to Bob.",
        "Carol", "Bob", 60,
    },
    {
        "Carol would gain 55 happiness units by sitting next to David.",
        "Carol", "David", 55,
    },
    {
        "David would gain 46 happiness units by sitting next to Alice.",
        "David", "Alice", 46,
    },
    {
        "David would lose 7 happiness units by sitting next to Bob.",
        "David", "Bob", -7,
    },
    {
        "David would gain 41 happiness units by sitting next to Carol.",
        "David", "Carol", 41,
    },
}

func TestParseLine( t *testing.T ) {
    for _,test := range Tests {
        a, b, dist := day13.ParseLine( test.Arg )
        
        testutil.AssertEq( t, a, test.a )
        testutil.AssertEq( t, b, test.b )
        testutil.AssertEq( t, dist, test.dist )
    }
}

func TestHappiestSeating( t *testing.T ) {
    routes := routes.NewRoutes( false )
    
    for _,test := range Tests {
        a, b, dist := day13.ParseLine( test.Arg )
        routes.AddRoute( a, b, dist )
    }
    
    testutil.AssertEq( t, day13.HappiestSeating( routes ), 330 )
}

func TestPart1( t *testing.T ) {
    routes := routes.NewRoutes( false )
    
    lines := util.LineChannel( Input_File )
    for line := range lines {
        routes.AddRoute( day13.ParseLine( line ) )
    }
    
    testutil.AssertEq( t, day13.HappiestSeating( routes ), 618 )
}
