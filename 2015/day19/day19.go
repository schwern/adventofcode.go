package day19

import(
    "bytes"
    "math"
    "regexp"
    "strings"
    "github.com/schwern/adventofcode.go/util"
    "github.com/schwern/adventofcode.go/mapWithDefault"
    mapset "github.com/deckarep/golang-set"
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

// A* algorithm from the Wikipedia pseudocode description.
func (self *Machine) LeastTransforms( start, goal string ) []string {
    // Nodes already evaluated.
    closedSet := mapset.NewSet()
    
    // Discovered nodes to be evaluated.
    openSet := mapset.NewSetWith(start)
    
    // For each node, which can it be most efficiently reached from.
    cameFrom := make( map[string]string )
    
    // Cost of going from the start to a node.
    // MaxInt by default.
    gScore := mapWithDefault.New( math.MaxInt32 )
    gScore.Set(start, 0)
    
    // Best guess at going from the start to a node.
    // MaxInt by default.
    fScore := mapWithDefault.New( math.MaxInt32 )
    fScore.Set(start, self.guessCost(start, goal))
    
    for openSet.Cardinality() != 0 {
        current := self.leastCost(openSet, fScore)
        if current == goal {
            return self.reconstructPath(cameFrom, current)
        }
        
        openSet.Remove(current)
        closedSet.Add(current)
        
        for neighbor := range self.Transforms(current) {
            if closedSet.Contains(neighbor) {
                // We already evaluated this.
                continue
            }
            openSet.Add(neighbor)
            
            tentativegScore := gScore.Get(current).(int) + 1
            if tentativegScore >= gScore.Get(neighbor).(int) {
                // It's not a better path
                continue
            }
            
            // Best path yet!
            cameFrom[neighbor] = current
            gScore.Set(neighbor, tentativegScore)
            fScore.Set(
                neighbor,
                gScore.Get(neighbor).(int) + self.guessCost(neighbor, goal),
            )
        }
    }
    
    // There is no path.
    return nil
}

func (self *Machine) guessCost( here, goal string ) int {
    switch {
        case here == goal:
            return 0
        case len(here) >= len(goal):
            return math.MaxInt32
        default:
            return len(goal) - len(here)
    }
}

func (self *Machine) leastCost( set mapset.Set, costs *mapWithDefault.Map ) string {
    minCost := math.MaxInt32
    var minNode string
    
    for item := range set.Iterator().C {
        cost := costs.Get(item.(string)).(int)
        if cost < minCost {
            minCost = cost
            minNode = item.(string)
        }
    }
    
    return minNode
}

func (self *Machine) reconstructPath( cameFrom map[string]string, current string ) []string {
    totalPath := []string{current}
    var ok bool
    for {
        current, ok = cameFrom[current]
        if !ok {
            break
        }
        totalPath = append(totalPath, current)
    }
    return totalPath
}
