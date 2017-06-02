package day16_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/2015/day16"
    "github.com/schwern/adventofcode.go/testutil"
)

func TestParseSue( t *testing.T ) {
    want := day16.NewSue(
        1,
        day16.Compounds{ "goldfish": 9, "cars": 0, "samoyeds": 9 },
    )
    
    have := day16.ParseSue(`Sue 1: goldfish: 9, cars: 0, samoyeds: 9`)
    testutil.AssertEq( t, have, want )
}
