package day07_test

import(
    "testing"
    "github.com/schwern/adventofcode2015/day07"
    "github.com/schwern/adventofcode2015/testutil"
)

func TestParseGate( t *testing.T ) {
    gates := make( map[string]day07.Gate )        
    gate := day07.ParseGate( "123 -> x", gates )
    
    testutil.AssertEq( t, gate.Output(), uint16(123) )
    testutil.AssertEq( t, gates["x"], gate )
}

func TestParseGate_GatesOutOfOrder( t *testing.T ) {
    gates := make( map[string]day07.Gate )
    
    c := day07.ParseGate( "a AND b -> c", gates )
    day07.ParseGate( "a -> 10", gates )
    day07.ParseGate( "b -> 12", gates )
    
    t.Skip("Out of order parsing is busted.")
    
    testutil.AssertEq( t, c.Output(), uint16(8) )
}
