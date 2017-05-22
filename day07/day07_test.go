package day07_test

import(
    "testing"
    "github.com/schwern/adventofcode2015/day07"
    "github.com/schwern/adventofcode2015/testutil"
)

func TestConstGate( t *testing.T ) {
    gate := day07.NewConstGate( "d", 10 )
    
    testutil.AssertEq( t, gate.ID(), "d" )
    testutil.AssertEq( t, gate.Output(), uint16(10) )
}

func TestUnaryGate( t *testing.T ) {
    const_gate := day07.NewConstGate( "a", 10 )
    
    gate := day07.NewUnaryGate( "b", "NOT", const_gate )
    
    testutil.AssertEq( t, gate.ID(), "b" )
    testutil.AssertEq( t, gate.Output(), uint16(65525) )
}

func TestBinaryGate( t *testing.T ) {
    in1 := day07.NewConstGate( "x", 30 )
    in2 := day07.NewConstGate( "y", 3 )
        
    tests := []struct{ op string; want uint16 }{
        { "AND", 2 },
        { "OR", 31 },
        { "LSHIFT", 240 },
        { "RSHIFT", 3 },
    }
    
    for _, test := range tests {
        id := "z"
        gate := day07.NewBinaryGate( id, test.op, in1, in2 )
        testutil.AssertEq( t, gate.ID(), id )
        testutil.AssertEq( t, gate.Output(), test.want )
    }
}

func TestParseGate( t *testing.T ) {
    gates := make( map[string]day07.Gate )        
    gate := day07.ParseGate( "123 -> x", gates )
    
    testutil.AssertEq( t, gate.Output(), uint16(123) )
    testutil.AssertEq( t, gates["x"], gate )
}
