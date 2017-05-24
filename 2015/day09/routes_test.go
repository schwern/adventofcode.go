package day09_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/2015/day09"
    "github.com/schwern/adventofcode.go/testutil"
)

func TestShortestRoute( t *testing.T ) {
    routes := day09.NewRoutes()
    
    routes.SetRoute( "London", "Dublin", 464 )
    routes.SetRoute( "London", "Belfast", 518 )
    routes.SetRoute( "Dublin", "Belfast", 141 )
    
    t.Skip("ShortestRoute not defined")
    //testutil.AssertEq( t, routes.ShortestRoute(), 605 )
}

func TestSetGetRoute( t *testing.T ) {
    routes := day09.NewRoutes()
    
    routes.SetRoute( "London", "Dublin", 464 )
    routes.SetRoute( "London", "Belfast", 518 )

    testutil.AssertEq( t, routes.GetRoute( "London", "Dublin" ), 464 )
    testutil.AssertEq( t, routes.GetRoute( "Dublin", "London" ), 464 )
    testutil.AssertEq( t, routes.GetRoute( "Belfast", "London" ), 518 )
    testutil.AssertEq( t, routes.GetRoute( "London", "Belfast" ), 518 )
}
