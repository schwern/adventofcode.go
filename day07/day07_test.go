package day07_test

import(
    "testing"
    "github.com/schwern/adventofcode2015/day07"
    "github.com/schwern/adventofcode2015/testutil"
)

func TestConstOp( t *testing.T ) {
    op := day07.NewConstOp( "d", 10 )
    
    testutil.AssertEq( t, op.ID, "d" )
    testutil.AssertEq( t, op.Output(), uint16(10) )
}

func TestUnaryOp( t *testing.T ) {
    const_op := day07.NewConstOp( "a", 10 )
    
    op := day07.NewUnaryOp( "b", day07.NOT, const_op )
    
    testutil.AssertEq( t, op.ID, "b" )
    testutil.AssertEq( t, op.Output(), uint16(65525) )
}

func TestBinaryOp( t *testing.T ) {
    in1 := day07.NewConstOp( "x", 30 )
    in2 := day07.NewConstOp( "y", 3 )
        
    tests := []struct{ op day07.OpType; want uint16 }{
        { day07.AND, 2 },
        { day07.OR, 31 },
        { day07.LSHIFT, 240 },
        { day07.RSHIFT, 3 },
    }
    
    for _, test := range tests {
        id := "z"
        op := day07.NewBinaryOp( id, test.op, in1, in2 )
        testutil.AssertEq( t, op.ID, id )
        testutil.AssertEq( t, op.Output(), test.want )
    }
}
