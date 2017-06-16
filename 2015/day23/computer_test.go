package day23_test

import(
    "testing"
    "github.com/stvp/assert"
    "github.com/schwern/adventofcode.go/2015/day23"
)

func TestRunInstructions( t *testing.T ) {
    comp := day23.NewComputer()
    
    instructions := []day23.Instruction{
        day23.Instruction{ Cmd: "inc", Args: []interface{}{"a"} },
        day23.Instruction{ Cmd: "jio", Args: []interface{}{"a", 2} },
        day23.Instruction{ Cmd: "tpl", Args: []interface{}{"a"} },
        day23.Instruction{ Cmd: "inc", Args: []interface{}{"a"} },
    }
    
    err := comp.RunInstructions( instructions )
    
    assert.Nil( t, err )
    assert.Equal( t, comp.Registers["a"], uint(2) )
    assert.Equal( t, comp.Registers["b"], uint(0) )
}

func TestRunInstructionsBad( t *testing.T ) {
    comp := day23.NewComputer()

    instructions := []day23.Instruction{
        day23.Instruction{ Cmd: "inc", Args: []interface{}{"a"} },
        day23.Instruction{ Cmd: "bad", Args: []interface{}{"a", 2} },
        day23.Instruction{ Cmd: "tpl", Args: []interface{}{"a"} },
        day23.Instruction{ Cmd: "inc", Args: []interface{}{"a"} },
    }
    
    err := comp.RunInstructions( instructions )

    assert.Equal( t, err.Error(), "Bad instruction" )
}
