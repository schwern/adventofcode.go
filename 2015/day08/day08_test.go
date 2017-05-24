package day08_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/util"
    "github.com/schwern/adventofcode.go/2015/day08"
)

var input_file = "testdata/input.txt"

func TestEncodingSize( t *testing.T ) {
    tests := []testutil.TestCase{
        { `""`, 6 },
        { `"abc"`, 9 },
        { `"aaa\"aaa"`, 16 },
        { `"\x27"`, 11 },
    }
    
    for _, test := range tests {
        testutil.AssertEq( t, day08.EncodingSize( test.Arg ), test.Want )
    }
}

func TestPart2( t *testing.T ) {
    lines := util.LineChannel( input_file )
    
    size := 0
    encoded_size := 0
    
    for line := range lines {
        size += len(line)
        encoded_size += day08.EncodingSize( line )
    }
    
    testutil.AssertEq( t, encoded_size - size, 2085 )
}
