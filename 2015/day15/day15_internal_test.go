package day15

import(
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
)

func TestIngredientComboChan( t *testing.T ) {
    var want = [][]int{
        {3,0,0},
        {2,1,0},
        {1,2,0},
        {0,3,0},
        {2,0,1},
        {1,1,1},
        {0,2,1},
        {1,0,2},
        {0,1,2},
        {0,0,3},
        nil,
    }
    
    c := NewCookie(3)
    for i := 0; i < 3; i++ {
        c.AddIngredient( Ingredient{} )
    }
    
    ch := c.ingredientComboChan()
    i := 0
    for combo := range ch {
        testutil.AssertEq( t, combo, want[i] )
        i++
    }
}
