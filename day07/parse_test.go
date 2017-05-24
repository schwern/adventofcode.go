package day07_test

import(
    "math/rand"
    "time"
    "testing"
    "github.com/schwern/adventofcode2015/day07"
    "github.com/schwern/adventofcode2015/testutil"
    "github.com/schwern/adventofcode2015/util"
)

var input_file = "../testdata/day07.txt"

func TestParseGate( t *testing.T ) {
    tests := []struct{ arg string; id string; op string; inputs []string }{
        { "123 -> x",       "x", "CONST", []string{"123"} },
        { "x AND y -> z",   "z", "AND",   []string{"x", "y"} },
        { "NOT e -> f",     "f", "NOT",   []string{"e"} },
    }
    
    for _, test := range tests {
        id, op, inputs := day07.ParseGate(test.arg)
    
        testutil.AssertEq( t, id, test.id )
        testutil.AssertEq( t, op, test.op )
        testutil.AssertStringSliceEq( t, inputs, test.inputs )
    }
}

func TestReadGates_InOrder( t *testing.T ) {
    lines := []string{
        "123 -> x",
        "456 -> y",
        "x AND y -> d",
        "x OR y -> e",
        "x LSHIFT 2 -> f",
        "y RSHIFT 2 -> g",
        "NOT x -> h",
        "NOT y -> i",
    }
    
    tests := []struct{ id string; want uint16 }{
        { "d", 72 },
        { "e", 507 },
        { "f", 492 },
        { "g", 114 },
        { "h", 65412 },
        { "i", 65079 },
        { "x", 123 },
        { "y", 456 },
    }
    
    testReadGates( t, lines, tests )
    shuffle(lines)
    testReadGates( t, lines, tests )
}

func testReadGates( t *testing.T, lines []string, tests []struct{id string;want uint16} ) {
    lineCh := make( chan string, len(lines) )
    for _, line := range lines {
        lineCh <- line
    }
    close(lineCh)

    gates := day07.ReadGates( lineCh )
    
    testutil.AssertEq( t, len(gates), len(lines) )

    for _, test := range tests {
        testutil.AssertEq( t, gates[test.id].Output(), test.want )
    }
}

func shuffle( list []string ) {
    rand.Seed( time.Now().Unix() )
    
    for i := range list {
        j := rand.Intn( i + 1 )
        list[i], list[j] = list[j], list[i]
    }
}

func TestDay07Part01( t *testing.T ) {
    gates := day07.ReadGates( util.LineChannel( input_file ) )
    testutil.AssertEq( t, gates["a"].Output(), uint16(16076) )
}
