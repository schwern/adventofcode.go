package main

import "testing"

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
        if have != test.want {
            t.Errorf(
                "FindFloor(%q) == %v, want %v",
                test.arg, have, test.want,
            )
        }
    }
}
