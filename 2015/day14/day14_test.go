package day14_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/2015/day14"
    "github.com/schwern/adventofcode.go/testutil"
)

func TestRunRunReindeer( t *testing.T ) {
    tests := []struct{ Arg day14.Reindeer; Want int }{
        {
            Arg: day14.Reindeer{ Speed: 14, Duration: 10, Rest: 127 },
            Want: 1120,
        },
        { 
            Arg: day14.Reindeer{ Speed: 16, Duration: 11, Rest: 162 },
            Want: 1056,
        },
    }
    
    for _,test := range tests {
        have := test.Arg.RunRunReindeer( 1000 )
        testutil.AssertEq( t, have, test.Want )
    }
}
