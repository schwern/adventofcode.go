package day23

import(
    "errors"
)

type Computer struct {
    Registers map[string]uint
    Pos int
}

func NewComputer() *Computer {
    self := Computer{
        Registers: make( map[string]uint ),
        Pos: 0,
    }
    
    return &self
}

func (self *Computer) RunInstruction( inst Instruction ) error {
    switch inst.Cmd {
        case "hlf":
            reg := inst.Args[0].(string)
            self.Registers[reg] /= 2
            self.Pos++
        case "tpl":
            reg := inst.Args[0].(string)
            self.Registers[reg] *= 3
            self.Pos++
        case "inc":
            reg := inst.Args[0].(string)
            self.Registers[reg]++
            self.Pos++
        case "jmp":
            self.Pos += inst.Args[0].(int)
        case "jie":
            reg := inst.Args[0].(string)
            if self.Registers[reg] % 2 == 0 {
                self.Pos += inst.Args[1].(int)
            }
        case "jio":
            reg := inst.Args[0].(string)
            if self.Registers[reg] == 1 {
                self.Pos += inst.Args[1].(int)
            }
        default:
            return errors.New("Bad instruction")
    }
    
    return nil
}

func (self *Computer) RunInstructions( insts []Instruction ) error {
    for self.Pos < len(insts) {
        err := self.RunInstruction( insts[self.Pos] )
        if err != nil {
            return err
        }
    }
    
    return nil
}