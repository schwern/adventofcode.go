package day23_test

import(
    "testing"
    "github.com/stvp/assert"
    "github.com/schwern/adventofcode.go/2015/day23"
    "github.com/schwern/adventofcode.go/util"
)

var InputFile = "testdata/input.txt"

func TestRunJio( t *testing.T ) {
    comp := day23.NewComputer()
    jio := day23.Instruction{ Cmd: "jio", Args: []interface{}{"a", 3} }
    
    tests := []struct{ reg uint; want int }{
        { 0, 1 },
        { 1, 3 },
        { 10, 1 },
    }
    
    for _,test := range tests {
        comp.Registers["a"] = test.reg
        comp.Pos = 0
        comp.RunInstruction( jio )
        assert.Equal( t, comp.Pos, test.want )
    }
}

func TestRunJie( t *testing.T ) {
    comp := day23.NewComputer()
    jie := day23.Instruction{ Cmd: "jie", Args: []interface{}{"a", 3} }
    
    tests := []struct{ reg uint; want int }{
        { 5, 1 },
        { 2, 3 },
        // 0 is even
        { 0, 3 },
    }
    
    for _,test := range tests {
        comp.Registers["a"] = test.reg
        comp.Pos = 0
        comp.RunInstruction( jie )
        assert.Equal( t, comp.Pos, test.want )
    }
}

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

func TestPart1( t *testing.T ) {
    comp := day23.NewComputer()
    
    insts := []day23.Instruction{}
    for line := range util.LineChannel( InputFile ) {
        insts = append( insts, day23.ParseInstruction(line) )
    }
    
    comp.RunInstructions(insts)
    assert.Equal( t, comp.Registers["b"], uint(184) )
}
