package day09_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/2015/day09"
    "github.com/schwern/adventofcode.go/testutil"
)

func TestAddGetRoute( t *testing.T ) {
    routes := day09.NewRoutes( true )
        
    routes.AddRoute( "London", "Dublin", 464 )
    routes.AddRoute( "London", "Belfast", 518 )
    routes.AddRoute( "Dublin", "Belfast", 141 )

    testutil.AssertEq( t, routes.GetRoute( "London", "Dublin" ), 464 )
    testutil.AssertEq( t, routes.GetRoute( "Dublin", "London" ), 464 )
    testutil.AssertEq( t, routes.GetRoute( "Belfast", "London" ), 518 )
    testutil.AssertEq( t, routes.GetRoute( "London", "Belfast" ), 518 )
    testutil.AssertEq( t, routes.GetRoute( "Dublin", "Belfast" ), 141 )
    testutil.AssertEq( t, routes.GetRoute( "Belfast", "Dublin" ), 141 )
}

func TestGetRouteUnknownRoute( t *testing.T ) {
    routes := day09.NewRoutes( false )

    defer testutil.AssertPanicf( t, "Do not have a node named Foo" )
    routes.GetRoute( "Foo", "Bar" )
}
