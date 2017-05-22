package day07

import(
    "regexp"
    "strconv"
    "github.com/schwern/adventofcode2015/util"
)

type Gate interface {
    Output() uint16
    ID() string
}

type BaseGate struct {
    id string
}

func (self *BaseGate) ID() string {
    return self.id
}

type ConstGate struct {
    BaseGate
    val uint16
}

func (self *ConstGate) Output() uint16 {
    return self.val
}

func NewConstGate( id string, val uint16 ) *ConstGate {
    self := ConstGate{ val: val }
    self.id = id
    
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

var gateRe = regexp.MustCompile(
    `(?:(?P<val>\w+)|(?P<in1>\w+)? (?P<op>\w+) (?P<in2>\w+)) -> (?P<id>\w+)`,
)
func ParseGate( line string, gates map[string]Gate ) Gate {
    gateInfo := util.FindAllNamed(gateRe, line)

    var gate Gate
    switch gateInfo["op"] {
        case "NOT":
            in := getGate( gateInfo["in"], gates )
            gate = NewUnaryGate( gateInfo["id"], gateInfo["op"], in )
        case "AND", "OR", "LSHIFT", "RSHIFT":
            in1 := getGate( gateInfo["in1"], gates )
            in2 := getGate( gateInfo["in2"], gates )
            gate = NewBinaryGate( gateInfo["id"], gateInfo["op"], in1, in2 )
        default:
            val, err := parseUint16( gateInfo["val"] )
            if err == nil {
                gate = NewConstGate( gateInfo["id"], val )
            } else {
                // Passthrough gate
                in := getGate( gateInfo["id"], gates ) 
                gate = NewUnaryGate( gateInfo["id"], "PASS", in )
            }
    }
    
    gates[gate.ID()] = gate
    
    return gate
}

func getGate( maybe string, gates map[string]Gate ) Gate {
    num, err := parseUint16( maybe )
    if err == nil {
        return NewConstGate( "", num )
    } else {
        return gates[maybe]
    }
}

func parseUint16( maybe string ) (uint16, error) {
    num, err := strconv.ParseUint( maybe, 10, 16 )
    return uint16(num), err
}
