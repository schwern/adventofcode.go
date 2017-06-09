package day19

import(
    "math"
    "testing"
    "github.com/stvp/assert"
    mapset "github.com/deckarep/golang-set"
    "github.com/schwern/adventofcode.go/mapWithDefault"
)

func TestReconstructPath( t *testing.T ) {
    machine := NewMachine()
    
    cameFrom := make( map[string]string )
    cameFrom["b"] = "a"
    cameFrom["c"] = "b"
    cameFrom["q"] = "z"
    
    assert.Equal( t, machine.reconstructPath(cameFrom, "c"), []string{"c", "b", "a"} )
    assert.Equal( t, machine.reconstructPath(cameFrom, "q"), []string{"q", "z"} )
    assert.Equal( t, machine.reconstructPath(cameFrom, "blah"), []string{"blah"} )
}

func TestLeastCost( t *testing.T ) {
    machine := NewMachine()
    
    costs := mapWithDefault.New( 9999 )
    costs.Set("foo", 23)
    costs.Set("bar", -5)
    
    set := mapset.NewSetWith("foo", "bar", "baz")
    assert.Equal( t, machine.leastCost( set, costs ), "bar" )
}

func TestGuessCost( t *testing.T ) {
    machine := NewMachine()
    
    assert.Equal( t, machine.guessCost( "123", "123456789" ), 6 )
    assert.Equal( t, machine.guessCost( "1234", "4321" ), math.MaxInt32 )
    assert.Equal( t, machine.guessCost( "12345", "4321" ), math.MaxInt32 )
}

