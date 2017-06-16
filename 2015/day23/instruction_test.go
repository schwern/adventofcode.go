package day23_test

import(
    "testing"
    "github.com/stvp/assert"
    "github.com/schwern/adventofcode.go/2015/day23"
)

func TestParseInstruction( t *testing.T ) {
    tests := []struct{ arg string; want day23.Instruction }{
        {
            arg: "inc a",
            want: day23.Instruction{
                Cmd: "inc",
                Args: []interface{}{"a"},
            },
        },
        {
            arg: "hlf r",
            want: day23.Instruction{
                Cmd: "hlf",
                Args: []interface{}{"r"},
            },
        },
        {
            arg: "tpl b",
            want: day23.Instruction{
                Cmd: "tpl",
                Args: []interface{}{"b"},
            },
        },
        {
            arg: "jmp -2",
            want: day23.Instruction{
                Cmd: "jmp",
                Args: []interface{}{-2},
            },
        },
        {
            arg: "jie a, +2",
            want: day23.Instruction{
                Cmd: "jie",
                Args: []interface{}{"a",2},
            },
        },
        {
            arg: "jio b, -3",
            want: day23.Instruction{
                Cmd: "jio",
                Args: []interface{}{"b",-3},
            },
        },
    }
    
    for _,test := range tests {
        have := day23.ParseInstruction(test.arg)
        assert.Equal( t, have, test.want )
    }
}

func TestParseInstructionBad( t *testing.T ) {
    defer assert.Panic(t, "Unknown instruction: foo")
    day23.ParseInstruction("foo b")
}
