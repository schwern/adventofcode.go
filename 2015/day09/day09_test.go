package day09_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/2015/day09"
    "github.com/schwern/adventofcode.go/testutil"
)

func TestBruteForce( t *testing.T ) {
    routes := day09.NewRoutes( true )
        
    routes.AddRoute( "London", "Dublin", 464 )
    routes.AddRoute( "London", "Belfast", 518 )
    routes.AddRoute( "Dublin", "Belfast", 141 )
    
    testutil.AssertEq( t, day09.BruteForce( routes ), 605 )
}

func TestParseLine( t *testing.T ) {
    a, b, dist := day09.ParseLine( "London to Dublin = 464" )
    
    testutil.AssertEq( t, a, "London" )
    testutil.AssertEq( t, b, "Dublin" )
    testutil.AssertEq( t, dist, 464 )
}
