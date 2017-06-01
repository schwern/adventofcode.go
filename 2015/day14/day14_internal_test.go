package day14

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
)

func TestIsRunning( t *testing.T ) {
    r := Reindeer{ Speed: 14, Duration: 10, Rest: 127 }
    
    tests := []struct{ Arg int; Want bool }{
        // These are all boundry changes.
        // { 11, false} implies { 10, true }
        { 0, true },
        { 10, false },
        { 137, true },
        { 147, false },
    }
        
    for _,test := range tests {        
        testutil.AssertEq( t, r.isRunning( test.Arg ), test.Want )
        testutil.AssertEq( t, r.isRunning( test.Arg-1 ), !test.Want )
    }
}
