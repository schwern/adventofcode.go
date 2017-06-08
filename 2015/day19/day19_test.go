package day19_test

import(
    "sort"
    "testing"
    "github.com/stvp/assert"
    "github.com/schwern/adventofcode.go/util"
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
        
        assert.Equal( t, have, want.val )
    }
        
    assert.Equal( t, machine.CountDistinctResults("HOH"), 4 )
    assert.Equal( t, machine.CountDistinctResults("HOHOHO"), 7 )
}

func TestTransform( t *testing.T ) {
    machine := day19.NewMachine()
    
    assert.Equal( t, machine.Transform("HOH", "H", "OH", 0), "OHOH" )    
    assert.Equal( t, machine.Transform("HOH", "H", "OH", 2), "HOOH" )    
    assert.Equal( t, machine.Transform("1Ca2", "Ca", "P", 1), "1P2" )
}

func TestPart1( t *testing.T ) {
    machine := day19.NewMachine()
    start := machine.ParseMachine( util.LineChannel( InputFile ) )

    assert.Equal( t, machine.CountDistinctResults(start), 518 )
}
