package day01

import "testing"
import "io/ioutil"

var input_file = "inputs/day01.txt"

type testCase struct {
    arg string
    want int
}

func assertEq(t *testing.T, have, want interface{}) {
    if have != want {
        t.Errorf(
            "\nhave: %v\nwant: %v",
            have, want,
        )
    }
}

func readInput( file string ) string {
    input, err := ioutil.ReadFile(file);
    if err != nil {
        panic(err)
    }
    
    return string(input)
}

func TestFindFloor(t *testing.T) {
    tests := []testCase {
        {"(())", 0},
        {"()()", 0},
        {"(((", 3},
        {"(()(()(", 3},
        {"))(((((", 3},
        {"())", -1},
        {"))(", -1},
        {")))", -3},
        {")())())", -3},
    }
    
    input := readInput(input_file)
    tests = append( tests, testCase{ input, 138 } )
    
    for _, test := range tests {
        have := FindFloor(test.arg)
        assertEq( t, have, test.want )
    }
}

func TestFirstBasement(t *testing.T) {
    tests := []testCase {
        { ")", 1 },
        { "()())", 5 },
        { "((", 0 },
    }
    
    input := readInput(input_file)
    tests = append( tests, testCase{ input, 1771 } )
    
    for _, test := range tests {
        have := FirstBasement(test.arg)
        assertEq( t, have, test.want )
    }
}
