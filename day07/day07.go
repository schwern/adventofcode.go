package day07

import(
    "regexp"
    "strconv"
    "github.com/schwern/adventofcode2015/util"
)

type Oper interface {
    Output() uint16
    ID() string
}

type BaseOp struct {
    id string
}

func (self *BaseOp) ID() string {
    return self.id
}

type ConstOp struct {
    BaseOp
    val uint16
}

func (self *ConstOp) Output() uint16 {
    return self.val
}

func NewConstOp( id string, val uint16 ) *ConstOp {
    self := new(ConstOp)
    self.id = id
    self.val = val
    
    return self
}

type UnaryOp struct {
    BaseOp
    op string
    in Oper
}

func NewUnaryOp( id string, op string, in Oper ) *UnaryOp {
    self := new(UnaryOp)
    self.id = id
    self.op = op
    self.in = in
    
    return self
}

func (self *UnaryOp) Output() uint16 {
    switch self.op {
        case "NOT":
            return ^self.in.Output()
        case "PASS":
            return self.in.Output()
        default:
            util.Panicf("Unknown op: %v", self.op)
            return 0
    }
}

type BinaryOp struct {
    BaseOp
    op string
    in1 Oper
    in2 Oper
}

func NewBinaryOp( id string, op string, in1 Oper, in2 Oper ) *BinaryOp {
    self := new(BinaryOp)
    self.id = id
    self.op = op
    self.in1 = in1
    self.in2 = in2
    
    return self
}

func (self *BinaryOp) Output() uint16 {
    in1 := self.in1.Output()
    in2 := self.in2.Output()
    
    switch self.op {
        case "AND":
            return in1 & in2
        case "OR":
            return in1 | in2
        case "LSHIFT":
            return in1 << in2
        case "RSHIFT":
            return in1 >> in2
        default:
            util.Panicf("Unknown op %v", self.op)
            return 0
    }
}

var opRe = regexp.MustCompile(
    `(?:(?P<val>\w+)|(?P<in1>\w+)? (?P<op>\w+) (?P<in2>\w+)) -> (?P<id>\w+)`,
)
func ParseOp( line string, ops map[string]Oper ) Oper {
    opInfo := util.FindAllNamed(opRe, line)

    var op Oper
    switch opInfo["op"] {
        case "NOT":
            in := getOp( opInfo["in"], ops )
            op = NewUnaryOp( opInfo["id"], opInfo["op"], in )
        case "AND", "OR", "LSHIFT", "RSHIFT":
            in1 := getOp( opInfo["in1"], ops )
            in2 := getOp( opInfo["in2"], ops )
            op = NewBinaryOp( opInfo["id"], opInfo["op"], in1, in2 )
        default:
            val, err := parseUint16( opInfo["val"] )
            if err == nil {
                op = NewConstOp( opInfo["id"], val )
            } else {
                // Passthrough op
                in := getOp( opInfo["id"], ops ) 
                op = NewUnaryOp( opInfo["id"], "PASS", in )
            }
    }
    
    ops[op.ID()] = op
    
    return op
}

func getOp( maybe string, ops map[string]Oper ) Oper {
    num, err := parseUint16( maybe )
    if err == nil {
        return NewConstOp( "", num )
    } else {
        return ops[maybe]
    }
}

func parseUint16( maybe string ) (uint16, error) {
    num, err := strconv.ParseUint( maybe, 10, 16 )
    return uint16(num), err
}
