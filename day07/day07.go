package day07

import(
    "github.com/schwern/adventofcode2015/util"
)

type OpType uint8

const(
    AND OpType = iota
    OR
    LSHIFT
    RSHIFT
    NOT
    CONST
)

type Oper interface {
    Output() uint16
}

type ConstOp struct {
    ID string
    val uint16
}

func (self *ConstOp) Output() uint16 {
    return self.val
}

func NewConstOp( id string, val uint16 ) *ConstOp {
    self := new(ConstOp)
    self.ID = id
    self.val = val
    
    return self
}

type UnaryOp struct {
    ID string
    op OpType
    in Oper
}

func NewUnaryOp( id string, op OpType, in Oper ) *UnaryOp {
    self := new(UnaryOp)
    self.ID = id
    self.op = op
    self.in = in
    
    return self
}

func (self *UnaryOp) Output() uint16 {
    return ^self.in.Output()
}

type BinaryOp struct {
    ID string
    op OpType
    in1 Oper
    in2 Oper
}

func NewBinaryOp( id string, op OpType, in1 Oper, in2 Oper ) *BinaryOp {
    self := new(BinaryOp)
    self.ID = id
    self.op = op
    self.in1 = in1
    self.in2 = in2
    
    return self
}

func (self *BinaryOp) Output() uint16 {
    in1 := self.in1.Output()
    in2 := self.in2.Output()
    
    switch self.op {
        case AND:
            return in1 & in2
        case OR:
            return in1 | in2
        case LSHIFT:
            return in1 << in2
        case RSHIFT:
            return in1 >> in2
        default:
            util.Panicf("Unknown op %v", self.op)
            return 0
    }
}
