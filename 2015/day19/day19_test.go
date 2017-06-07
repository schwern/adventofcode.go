package day19_test

import(
    "sort"
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/2015/day19"
)

var InputFile = "testdata/input.txt"

func TestAddTransformString( t *testing.T ) {
    haves := []string {
        "H => HO",
        "H => OH",
        "O => HH",
        "Ca => PB",
    }
    
    wants := []struct{ key string; val []string }{
        { "H", []string{ "HO", "OH" } },
        { "O", []string{ "HH" } },
        { "Ca", []string{ "PB" } },
    }
    
    machine := day19.NewMachine()
    for _,have := range haves {
        machine.AddTransformString( have )
    }
    
    for _,want := range wants {
        have := machine.GetTransforms(want.key)

        sort.Strings( have )
        sort.Strings( want.val )        
        
        testutil.AssertEq( t, have, want.val )
    }
        
    testutil.AssertEq( t, machine.CountDistinctResults("HOH"), 4 )
    testutil.AssertEq( t, machine.CountDistinctResults("HOHOHO"), 7 )
}

func TestTransform( t *testing.T ) {
    machine := day19.NewMachine()
    
    testutil.AssertEq( t, machine.Transform("HOH", "H", "OH", 0), "OHOH" )    
    testutil.AssertEq( t, machine.Transform("HOH", "H", "OH", 2), "HOOH" )    
    testutil.AssertEq( t, machine.Transform("1Ca2", "Ca", "P", 1), "1P2" )
}
