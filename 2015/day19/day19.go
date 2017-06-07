package day19

import(
    "regexp"
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

func (self *Machine) ParseTransforms( lines chan string ) {
    for line := range lines {
        self.AddTransformString( line )
    }
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
