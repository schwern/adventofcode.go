package routes_test

import(
    "sort"
    "testing"
    "github.com/schwern/adventofcode.go/routes"
    "github.com/schwern/adventofcode.go/testutil"
)

func TestAddGetRoute( t *testing.T ) {
    routes := routes.NewRoutes( true )
        
    routes.AddRoute( "London", "Dublin", 464 )
    routes.AddRoute( "London", "Belfast", 518 )
    routes.AddRoute( "Dublin", "Belfast", 141 )

    testutil.AssertEq( t, routes.MustGetRoute( "London", "Dublin" ), 464 )
    testutil.AssertEq( t, routes.MustGetRoute( "Dublin", "London" ), 464 )
    testutil.AssertEq( t, routes.MustGetRoute( "Belfast", "London" ), 518 )
    testutil.AssertEq( t, routes.MustGetRoute( "London", "Belfast" ), 518 )
    testutil.AssertEq( t, routes.MustGetRoute( "Dublin", "Belfast" ), 141 )
    testutil.AssertEq( t, routes.MustGetRoute( "Belfast", "Dublin" ), 141 )
}

func TestGetRouteUnknownRoute( t *testing.T ) {
    routes := routes.NewRoutes( false )

    defer testutil.AssertPanicf( t, "Do not have a node named Foo" )
    routes.GetRoute( "Foo", "Bar" )
}

func TestTryAllPaths( t *testing.T ) {
    routes := routes.NewRoutes( true )
        
    routes.AddRoute( "London", "Dublin", 464 )
    routes.AddRoute( "London", "Belfast", 518 )
    routes.AddRoute( "Dublin", "Belfast", 141 )
    
    ch := routes.TryAllPaths()
    have := []int{}
    for cost := range ch {
        have = append(have, cost)
    }
    sort.Ints(have)
    want := []int{ 605, 605, 659, 659, 982, 982 }
    
    testutil.AssertIntSliceEq( t, have, want )
}
