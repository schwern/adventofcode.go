package day08_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/util"
    "github.com/schwern/adventofcode.go/2015/day08"
)

var input_file = "testdata/input.txt"

func TestMemorySize( t *testing.T ) {
    tests := []testutil.TestCase{
        { `""`, 0 },
        { `"abc"`, 3 },
        { `"aaa\"aaa"`, 7 },
        { `"\x27"`, 1 },
    }
    
    for _, test := range tests {
        testutil.AssertEq( t, day08.MemorySize( test.Arg ), test.Want )
    }
}

func TestPart1( t *testing.T ) {
    lines := util.LineChannel( input_file )
    
    size := 0
    mem_size := 0
    
    for line := range lines {
        size += len(line)
        mem_size += day08.MemorySize( line )
    }
    
    testutil.AssertEq( t, size - mem_size, 1350 )
}

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
