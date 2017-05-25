package day09_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/2015/day09"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/util"
)

var input_file = "testdata/input.txt"

func TestBruteForce( t *testing.T ) {
    routes := day09.NewRoutes( true )
        
    routes.AddRoute( "London", "Dublin", 464 )
    routes.AddRoute( "London", "Belfast", 518 )
    routes.AddRoute( "Dublin", "Belfast", 141 )
    
    testutil.AssertEq( t, day09.BestRouteBruteForce( routes ), 605 )
    testutil.AssertEq( t, day09.WorstRouteBruteForce( routes ), 982 )
}

func TestParseLine( t *testing.T ) {
    a, b, dist := day09.ParseLine( "London to Dublin = 464" )
    
    testutil.AssertEq( t, a, "London" )
    testutil.AssertEq( t, b, "Dublin" )
    testutil.AssertEq( t, dist, 464 )
}

func TestPart1( t *testing.T ) {
    lines := util.LineChannel( input_file )
    routes := day09.NewRoutes( true )
    
    for line := range lines {
        a, b, dist := day09.ParseLine( line )
        routes.AddRoute( a, b, dist )
    }
    
    testutil.AssertEq( t, day09.BestRouteBruteForce( routes ), 251 )
}

func TestPart2( t *testing.T ) {
    lines := util.LineChannel( input_file )
    routes := day09.NewRoutes( true )
    
    for line := range lines {
        a, b, dist := day09.ParseLine( line )
        routes.AddRoute( a, b, dist )
    }
    
    testutil.AssertEq( t, day09.WorstRouteBruteForce( routes ), 898 )
}
