package day09_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/2015/day09"
    "github.com/schwern/adventofcode.go/testutil"
)

func TestNewHeldKarp( t *testing.T ) {
    hk := day09.NewHeldKarp(nil)
    hk.AddRoute( "London", "Dublin", 464 )
    hk.AddRoute( "London", "Belfast", 518 )
    hk.AddRoute( "Dublin", "Belfast", 141 )
    
    testutil.AssertEq( t, hk.GetRoute( "London", "Dublin" ), 464 )
    // Routes are not symetrical by default.
    testutil.AssertEq( t, hk.GetRoute( "Dublin", "London" ), 0 )
}

func TestNewHeldKarpWithRoutes( t *testing.T ) {
    routes := day09.NewRoutes( true )
    routes.AddRoute( "London", "Dublin", 464 )

    hk := day09.NewHeldKarp( routes )
    hk.AddRoute( "Dublin", "Belfast", 141 )

    // HK uses the existing routes
    testutil.AssertEq( t,
        hk.GetRoute("London", "Dublin"),
        routes.GetRoute("London", "Dublin"),
    )

    // HK holds a pointer to its Routes
    testutil.AssertEq( t,
        hk.GetRoute("Dublin", "Belfast"),
        routes.GetRoute("Dublin", "Belfast" ),
    )
}
