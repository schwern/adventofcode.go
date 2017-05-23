package day07

import(
    "regexp"
    "strconv"
    "github.com/schwern/adventofcode2015/util"
)

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
