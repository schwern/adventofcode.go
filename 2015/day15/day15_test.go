package day15_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/util"
    "github.com/schwern/adventofcode.go/2015/day15"
)

var InputFile = "testdata/input.txt"

var Butterscotch = day15.Ingredient{
    Capacity: -1,
    Durability: -2,
    Flavor: 6,
    Texture: 3,
    Calories: 8,
}
var Cinnamon = day15.Ingredient{
    Capacity: 2,
    Durability: 3,
    Flavor: -2,
    Texture: -1,
    Calories: 3,
}

func TestParseIngredient( t *testing.T ) {
    tests := []struct{ Arg string; Want day15.Ingredient }{
        {
            Arg: `Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8`,
            Want: Butterscotch,
        },
        {
            Arg: `Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`,
            Want: Cinnamon,
        },
    }
    
    for _,test := range tests {
        have := day15.ParseIngredient( test.Arg )
        testutil.AssertEq( t, have, test.Want )
    }
}

func TestCookieScore( t *testing.T ) {
    c := day15.NewCookie()
    c.AddIngredient( Butterscotch )
    c.AddIngredient( Cinnamon )
    
    testutil.AssertEq( t, c.Score([]int{44,56}), 62842880 )
    testutil.AssertEq( t, c.BestScore(100), 62842880 )
}

func TestPart1( t *testing.T ) {    
    c := day15.NewCookie()
    
    lines := util.LineChannel( InputFile )    
    for line := range lines {
        c.AddIngredient( day15.ParseIngredient(line) )
    }
    
    testutil.AssertEq( t, c.BestScore(100), 13882464 )
}

func TestPart2( t *testing.T ) {    
    c := day15.NewCookie()
    
    lines := util.LineChannel( InputFile )    
    for line := range lines {
        c.AddIngredient( day15.ParseIngredient(line) )
    }
    
    testutil.AssertEq( t, c.BestScoreExactCalories(100, 500), 11171160 )
}