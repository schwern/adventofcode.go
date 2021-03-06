package day16_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/2015/day16"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/util"
)

var InputFile = "testdata/input.txt"

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

func TestPart1( t *testing.T ) {
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
    
    sueNum := 0
    for line := range util.LineChannel( InputFile ) {
        sue := day16.ParseSue( line )
        if sue.CheckCompounds(want) {
            sueNum = sue.Num
            break
        }
    }
    
    testutil.AssertEq( t, sueNum, 40 )
}

func TestPart2( t *testing.T ) {
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
    
    greaterThan := func(a int, b int) bool { return a > b }
    lessThan := func(a int, b int) bool { return a < b }
    
    ranges := map[string]day16.CompoundCheck{
        "cats": greaterThan,
        "trees": greaterThan,
        "pomeranians": lessThan,
        "goldfish": lessThan,
    }
    
    sueNum := 0
    for line := range util.LineChannel( InputFile ) {
        sue := day16.ParseSue( line )
        if sue.CheckCompoundsRanged(want, ranges) {
            sueNum = sue.Num
            break
        }
    }
    
    testutil.AssertEq( t, sueNum, 241 )
}
