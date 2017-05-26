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

func TestIncString( t *testing.T ) {
    tests := []string{
        "xx", "xy", "xz", "ya", "yb",
    }
    
    for i := 1; i < len(tests); i++ {
        testutil.AssertEq( t, day11.IncString(tests[i-1]), tests[i] )
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
