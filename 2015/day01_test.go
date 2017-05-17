package main

import "testing"
import "io/ioutil"

func assertEq(t *testing.T, have, want interface{}) {
    if have != want {
        t.Errorf(
            "\nhave: %v\nwant: %v",
            have, want,
        )
    }
}

func TestFindFloor(t *testing.T) {
    tests := []struct { arg string; want int } {
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
    
    for _, test := range tests {
        have := FindFloor(test.arg)
        assertEq( t, have, test.want )
    }
}

func TestAnswer(t *testing.T) {
    input, err := ioutil.ReadFile("inputs/day01.txt");
    if err != nil {
        panic(err)
    }
    
    have := FindFloor( string(input) )
    assertEq( t, have, 138 )
}