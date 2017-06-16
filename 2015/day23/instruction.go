package day23

import(
    "strings"
    "github.com/schwern/adventofcode.go/util"
)

type Instruction struct {
    Cmd string
    Args []interface{}
}

func NewInstruction() Instruction {
    return Instruction{
        Cmd: "",
        Args: make( []interface{}, 0, 2 ),
    }
}

func ParseInstruction( line string ) Instruction {
    inst := NewInstruction()
    
    pair := strings.SplitN( line, " ", 2 )
    inst.Cmd = pair[0]
    
    switch inst.Cmd {
        // cmd reg
        case "hlf", "tpl", "inc":
            inst.Args = append( inst.Args, pair[1] )
        // cmd int
        case "jmp":
            inst.Args = append( inst.Args, util.MustAtoi(pair[1]) )
        // cmd reg int
        case "jie", "jio":
            args := strings.Split( pair[1], ", " )
            inst.Args = append( inst.Args, args[0], util.MustAtoi(args[1]) )
        default:
            util.Panicf("Unknown instruction: %v", inst.Cmd)
    }
    
    return inst
}
