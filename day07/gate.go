package day07

import(
    "github.com/schwern/adventofcode2015/util"
)

type Gate interface {
    Output() uint16
    ID() string
    SetID(string)
    ClearOutputCache()
    SetOutputCache( uint16 )
}

type BaseGate struct {
    id string
    out uint16
    cached bool
}

func (self *BaseGate) ID() string {
    return self.id
}

func (self *BaseGate) SetID( id string ) {
    self.id = id
}

func (self *BaseGate) ClearOutputCache() {
    self.cached = false
}

func (self *BaseGate) SetOutputCache( val uint16 ) {
    self.out = val
    self.cached = true
}

type ConstGate struct {
    BaseGate
}

func (self *ConstGate) Output() uint16 {
    return self.out
}

func NewConstGate( id string, val uint16 ) *ConstGate {
    self := ConstGate{}
    self.id = id
    self.SetOutputCache( val )
    
    return &self
}

type UnaryGate struct {
    BaseGate
    op string
    in Gate
}

func NewUnaryGate( id string, op string, in Gate ) *UnaryGate {
    self := UnaryGate{ op: op, in: in }
    self.id = id

    return &self
}

func (self *UnaryGate) Output() uint16 {
    if self.cached {
        return self.out
    }
    
    var val uint16
    switch self.op {
        case "NOT":
            val = ^self.in.Output()
        case "PASS":
            val = self.in.Output()
        default:
            util.Panicf("Unknown op: %v", self.op)
            return 0
    }
    
    self.SetOutputCache( val )
    return val
}

type BinaryGate struct {
    BaseGate
    op string
    in1 Gate
    in2 Gate
}

func NewBinaryGate( ident string, op string, in1 Gate, in2 Gate ) *BinaryGate {
    self := BinaryGate{ op: op, in1: in1, in2: in2 }
    self.id = ident
    
    return &self
}

func (self *BinaryGate) Output() uint16 {
    if self.cached {
        return self.out
    }
    
    in1 := self.in1.Output()
    in2 := self.in2.Output()
    
    var val uint16
    switch self.op {
        case "AND":
            val = in1 & in2
        case "OR":
            val = in1 | in2
        case "LSHIFT":
            val = in1 << in2
        case "RSHIFT":
            val = in1 >> in2
        default:
            util.Panicf("Unknown op %v", self.op)
            return 0
    }
    
    self.SetOutputCache( val )
    return val
}

func MakeGate(id string, op string, inputs []Gate) Gate {
    if op == "CONST" {
        gate := inputs[0]
        gate.SetID(id)
        return gate
    }
    
    switch len(inputs) {
        case 1:
            return NewUnaryGate( id, op, inputs[0] )
        case 2:
            return NewBinaryGate( id, op, inputs[0], inputs[1] )
        default:
            util.Panicf("Wrong number of inputs: %v.", len(inputs))
            return nil
    }
}
