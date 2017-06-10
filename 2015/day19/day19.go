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

func (self *Machine) Transforms( start string ) chan string {
    out := make( chan string )
    
    go func() {
        defer close(out)
        
        for idx := range start {
            for from,list := range self.transforms {
                str := start[idx:]
                if strings.HasPrefix( str, from ) {
                    for _,to := range list {
                        out <- self.Transform( start, from, to, idx )
                    }
                }
            }
        }
    }()
    
    return out
}

func (self *Machine) CountDistinctResults( start string ) int {
    results := make( map[string]bool )
    for new := range self.Transforms( start ) {
        results[new] = true
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

var numElementsRe = regexp.MustCompile(
    `([A-Z][a-z]|[A-Z]|e)`,
)
func (self *Machine) NumElements( molecule string ) int {
    if len(molecule) == 0 {
        return 0
    }
    
    matches := numElementsRe.FindAllString(molecule, -1)
    if matches == nil {
        util.Panicf("Cannot match %v against %v", numElementsRe, molecule)
    }
    
    return len(matches)
}

func (self *Machine) NumTransforms( start, goal string ) int {
    // All transforms are of these forms...
    // Note: y can be x, and two y's can be the same element,
    // but x and y are never Rn, Ar, or Y.
    // x => y y
    // x => y Rn y Ar
    // x => y Rn y Y y (Y y)* Ar
    //
    // Rn and Ar are always balanced.
    // Transforms always add 1 + the number of Y's.
    
    // x => yy always adds 1, so the number necessary
    // is the difference in length
    changes := self.NumElements(goal) - self.NumElements(start)
    
    // Special case for 'e' which might be e => X.
    // That adds an extra transform for each e.
    if len(self.GetTransforms("e")[0]) == 1 {
        eCount := strings.Count(start, "e")
        changes += eCount
    }
    
    // Rn/Ar pairs don't add an extra transform.
    rnCount := strings.Count(goal, "Rn")
    arCount := strings.Count(goal, "Ar")
    if rnCount != arCount {
        util.Panicf("rnCount: %v, arCount: %v. Should be equal!", rnCount, arCount)
    }    
    changes -= rnCount * 2
    
    // Each Y indicates the transform can do an extra one.
    yCount := strings.Count(goal, "Y")
    changes -= yCount * 2
    
    return changes
}
