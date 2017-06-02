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

func TestCheckCompounds( t *testing.T ) {
    want := day16.Compounds{
        "children": 3,
        "cats": 7,
        "samoyeds": 2,
        "pomeranians": 3,
        "akitas": 0,
        "vizslas": 0,
        "goldfish": 5,
        "trees": 3,
        "cars": 2,
        "perfumes": 1,
    }
    
    match := day16.NewSue(
        10,
        day16.Compounds{ "children": 3, "cats": 7, "cars": 2 },
    )
    testutil.AssertEq( t, match.CheckCompounds(want), true )
    
    notMatch := day16.NewSue(
        20,
        day16.Compounds{ "cats": 0 },
    )
    testutil.AssertEq( t, notMatch.CheckCompounds(want), false )
}
