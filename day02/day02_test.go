package day02_test

import(
    "testing"
    "strings"
    "strconv"
    "fmt"
    "github.com/schwern/adventofcode2015/testutil"
    "github.com/schwern/adventofcode2015/day02"
    "github.com/schwern/adventofcode2015/util"
)

var Input_File = "../testdata/day02.txt"

func TestWrappingPaperNeeded( t *testing.T ) {
    tests := []struct{ l, w, h int; want int } {
        { 2, 3, 4, 58 },
        { 1, 1, 10, 43 },
    }
    
    for _, test := range tests {
        have := day02.WrappingPaperNeeded( test.l, test.w, test.h )
        testutil.AssertEq( t, have, test.want )
    }
}

func TestRibbonNeeded( t *testing.T ) {
    tests := []struct{ l, w, h int; want int } {
        { 2, 3, 4, 34 },
        { 1, 1, 10, 14 },
    }
    
    for _, test := range tests {
        have := day02.RibbonNeeded( test.l, test.w, test.h )
        testutil.AssertEq( t, have, test.want )
    }
}

func parseInputs( filename string ) chan []int {    
    ch := make( chan []int )
    go func() {
        defer close(ch)
        lines := util.LineChannel(filename)

        for line := range lines {
            args := make( []int, 3 )
            for i, arg := range strings.Split(line, "x") {
                num, err := strconv.Atoi(arg)
                if err != nil {
                    panic(fmt.Sprintf("Can't convert %s to integer", arg))
                }
                
                args[i] = num
            }
            
            ch <- args
        }
    }()
    
    return ch
}

func TestDay02Part01( t *testing.T ) {    
    paper := 0

    for args := range parseInputs(Input_File) {
        paper += day02.WrappingPaperNeeded( args[0], args[1], args[2] )
    }
    
    testutil.AssertEq( t, paper, 1606483 )
}

func TestDay02Part02( t *testing.T ) {
    ribbon := 0

    for args := range parseInputs(Input_File) {
        ribbon += day02.RibbonNeeded( args[0], args[1], args[2] )
    }
    
    testutil.AssertEq( t, ribbon, 3842356 )
}
