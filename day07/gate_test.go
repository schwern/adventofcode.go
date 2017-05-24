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

func TestMakeGate( t *testing.T ) {
    in1 := day07.NewConstGate("", 123)
    in2 := day07.NewConstGate("", 234)
    
    tests := []struct{ id string; op string; inputs []day07.Gate; output uint16 }{
        { "x", "CONST", []day07.Gate{ in1 }, in1.Output() },
        { "z", "AND",   []day07.Gate{ in1, in2 }, in1.Output() & in2.Output() },
        { "f", "NOT",   []day07.Gate{ in1 }, ^in1.Output() },
    }
    
    for _, test := range tests {
        gate := day07.MakeGate( test.id, test.op, test.inputs )
        
        testutil.AssertEq( t, gate.ID(), test.id )
        testutil.AssertEq( t, gate.Output(), test.output )
    }
}