package day25_test

import(
    "testing"
    "github.com/stvp/assert"
    "github.com/schwern/adventofcode.go/2015/day25"
)

var Row = 2947
var Col = 3029
var StartCode = 20151125

func TestNextCode( t *testing.T ) {
    want := []int{
        StartCode,
        31916031, 18749137, 16080970, 21629792, 17289845,
        24592653, 8057251, 16929656, 30943339,
        77061, 32451966, 1601130, 7726640, 10071777,
        33071741, 17552253, 21345942, 7981243, 15514188, 33511524,
    }
    
    for i := 1; i < len(want); i++ {
        assert.Equal( t, day25.NextCode( want[i-1] ), want[i] )
    }
}

func TestGridPos( t *testing.T ) {
    tests := []struct{ row,col,want int }{
        {1,1,1},
        {1,2,3},
        {2,2,5},
        {5,2,17},
        {3,4,19},
    }
        
    for _,test := range tests {
        assert.Equal( t, day25.GridPos(test.row,test.col), test.want )
    }
}

func TestPart1( t *testing.T ) {
    gridNum := day25.GridPos(Row,Col)
    
    code := 20151125
    for i := 2; i <= gridNum; i++ {
        code = day25.NextCode(code)
    }
    
    assert.Equal( t, code, 19980801 )
}
