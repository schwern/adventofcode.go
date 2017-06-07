package day19

import(
    "bytes"
    "regexp"
    "strings"
    "github.com/schwern/adventofcode.go/util"
)

type Machine struct {
    transforms map[string][]string
}

func NewMachine() *Machine {
    self := Machine{
        transforms: make( map[string][]string ),
    }
    return &self
}

var transformRe = regexp.MustCompile( `(\w+) => (\w+)` )
func ( self *Machine ) parseLine( line string ) (string, string) {
    match := transformRe.FindStringSubmatch( line )
    if match == nil {
        util.Panicf( "Can't parse: %v", line )
    }
    
    return match[1], match[2]
}

func (self *Machine) ParseMachine( lines chan string ) string {
    for line := range lines {
        if line == "" {
            break
        }
        self.AddTransformString( line )
    }
    
    return <-lines
}

func (self *Machine) AddTransformString( line string ) {
    self.AddTransform( self.parseLine(line) )
}

func (self *Machine) AddTransform( key, val string ) {
    if list, ok := self.transforms[key]; ok {
        self.transforms[key] = append(list, val)
    } else {
        self.transforms[key] = []string{val}
    }
}

func (self *Machine) GetTransforms( key string ) []string {
    return self.transforms[key]
}

func (self *Machine) CountDistinctResults( start string ) int {
    results := make( map[string]bool )
    
    for idx := range start {
        for from,list := range self.transforms {
            str := start[idx:]
            if strings.HasPrefix( str, from ) {
                for _,to := range list {
                    new := self.Transform( start, from, to, idx )
                    results[new] = true
                }
            }
        }
    }
    
    return len(results)
}

func (self *Machine) Transform( start, from, to string, idx int ) string {
    var new bytes.Buffer
    new.Grow( len(start) + len(to) - len(from) )
    new.WriteString(start[:idx])
    new.WriteString(to)
    new.WriteString(start[idx+len(from):])
    
    return new.String()
}
