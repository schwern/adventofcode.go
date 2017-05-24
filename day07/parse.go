package day07

import(
    "regexp"
    "strconv"
    "sync"
    "github.com/schwern/adventofcode2015/util"
)

type Gates map[string]Gate

type readMsg struct {
    id string
    resp chan Gate
}

type writeMsg struct {
    gate Gate
}

func ReadGates( lines chan string ) Gates {
    gates, readCh, writeCh, doneCh := gatesKeeper()
    var wg sync.WaitGroup

    for line := range lines {
        wg.Add(1)
        go func( line string ){
            defer wg.Done()
            id, op, inputIDs := ParseGate( line )
            
            inputs := make([]Gate, len(inputIDs))
            for i := range inputIDs {
                inputs[i] = readGate( inputIDs[i], readCh )
            }
            
            gate := MakeGate( id, op, inputs )
            writeGate( gate, writeCh )
        }(line)
    }
    wg.Wait()
    doneCh<-true
    
    return gates
}

func readGate( maybe string, readCh chan *readMsg ) Gate {
    num, err := parseUint16( maybe )
    if err == nil {
        return NewConstGate( "", num )
    }
    
    respCh := make( chan Gate )
    msg := readMsg{ id: maybe, resp: respCh }
    readCh<-&msg
    return <-respCh
}

func writeGate( gate Gate, writeCh chan *writeMsg ) {
    msg := writeMsg{ gate: gate }
    writeCh <- &msg
}

func gatesKeeper() (Gates, chan *readMsg, chan *writeMsg, chan bool) {
    readCh  := make( chan *readMsg )
    writeCh := make( chan *writeMsg )
    doneCh  := make( chan bool )
    gates   := make( Gates )

    waiting := make( map[string][]chan Gate )
    
    go func() {
        for {            
            select {
                case read := <-readCh:
                    id := read.id
                    gate := gates[id]
                    if gate == nil {
                        if waiting[id] == nil {
                            waiting[id] = make( []chan Gate, 0, 1 )
                        }
                        waiting[id] = append(waiting[id], read.resp)
                    } else {
                        read.resp <- gate
                    }
                    
                case write := <-writeCh:
                    gate := write.gate
                    id := gate.ID()
                    gates[id] = gate

                    for _, ch := range waiting[id] {
                        ch <- gate
                    }
                    delete(waiting, id)
                    
                case <-doneCh:
                    return
            }
        }
    }()

    return gates, readCh, writeCh, doneCh
}

var gateRe = regexp.MustCompile(
    `^\s*(?:(?P<val>\w+)|(?:(?P<in1>\w+) )?(?P<op>\w+) (?P<in2>\w+)) -> (?P<id>\w+)\s*$`,
)
func ParseGate( line string ) (string, string, []string) {
    def := util.FindAllNamed(gateRe, line)

    id := def["id"]
    op := def["op"]
    inputs := make([]string, 1)

    switch op {
        case "NOT":
            inputs[0] = def["in2"]
        case "AND", "OR", "LSHIFT", "RSHIFT":
            inputs = []string{def["in1"], def["in2"]}
        default:
            op = "CONST"
            inputs[0] = def["val"]
    }

    return id, op, inputs
}

func parseUint16( maybe string ) (uint16, error) {
    num, err := strconv.ParseUint( maybe, 10, 16 )
    return uint16(num), err
}
