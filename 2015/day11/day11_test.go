package day11_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/2015/day11"
)

var input = "cqjxjnds"

func TestIsValidPassword( t *testing.T ) {
    tests := []struct{ arg string; want bool }{
        { "hijklmmn", false },
        { "abbceffg", false },
        { "abbcegjk", false },
        
        { "abcdffaa", true },
        // too short
        { "abcffaa", false },
        // too long
        { "abcdffaab", false },
        // invalid letter
        { "abclffaa", false },
        // invalid number
        { "abc1ffaa", false },
        // overlapping pair
        { "abcdfffa", false },
        { "ghjaabcc", true },
    }
    
    for _,test := range tests {
        have := day11.IsValidPassword( test.arg )
        testutil.AssertEq( t, have, test.want )
    }
}

func TestIncByteSlice( t *testing.T ) {
    tests := []string{
        "xx", "xy", "xz", "ya", "yb",
    }
    
    for i := 1; i < len(tests); i++ {
        have := []byte(tests[i-1])
        day11.IncByteSlice(have)
        testutil.AssertEq( t, string(have), tests[i] )
    }
}

func TestNextPassword( t *testing.T ) {
    tests := []struct{ arg string; want string }{
        { "abcdefgh", "abcdffaa" },
        { "ghijklmn", "ghjaabcc" },
    }
    
    for _,test := range tests {        
        have := day11.NextPassword( test.arg )
        testutil.AssertEq( t, have, test.want )
    }
}

func TestParts1And2( t *testing.T ) {
    part1 := day11.NextPassword( input )
    testutil.AssertEq( t, part1, "cqjxxyzz" )
    
    part2 := day11.NextPassword( part1 )
    testutil.AssertEq( t, part2, "cqkaabcc" )
}
