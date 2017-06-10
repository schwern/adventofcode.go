package day19

import(
    "bytes"
    "math"
    "regexp"
    "strings"
    "github.com/schwern/adventofcode.go/util"
    "github.com/schwern/adventofcode.go/mapWithDefault"
    "github.com/ferhatelmas/levenshtein"
    mapset "github.com/deckarep/golang-set"
    priorityQueue "github.com/tracymacding/priority-queue"
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
    
    // Queue of nodes to be evaluated.
    openSet := priorityQueue.NewPriorityQueue()
    
    // Add the start to the queue.
    // Priority is the estimate of its cost to reach the goal.
    openSet.Push(
        priorityQueue.NewNode(start, nil, self.guessCost(start, goal)),
    )
    
    // For each node, which can it be most efficiently reached from.
    cameFrom := make( map[string]string )
    
    // Cost of going from the start to a node.
    gScore := mapWithDefault.New( math.MaxInt32 )
    gScore.Set(start, 0)
    
    for openSet.Length() != 0 {
        currentNode := openSet.Pop()
        current := currentNode.GetKey().(string)
        if current == goal {
            return self.reconstructPath(cameFrom, current)
        }        
        closedSet.Add(current)
        
        for neighbor := range self.Transforms(current) {            
            if len(neighbor) > len(goal) {
                // It's too big
                continue
            }
            if len(neighbor) == len(goal) && neighbor != goal {
                // Right size, but not the goal
                continue
            }
            if closedSet.Contains(neighbor) {
                // We already evaluated this.
                continue
            }
            neighborNode := priorityQueue.NewNode(neighbor, nil, math.MaxInt32)
            
            tentativegScore := gScore.Get(current).(int) + 1
            if tentativegScore >= gScore.Get(neighbor).(int) {
                // It's not a better path
                openSet.Push( neighborNode )
                continue
            }
            
            // Best path yet!
            cameFrom[neighbor] = current
            gScore.Set(neighbor, tentativegScore)
            costGuess := gScore.Get(neighbor).(int) + self.guessCost(neighbor, goal)
            neighborNode.UpdatePrio( costGuess )
            openSet.Push( neighborNode )
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
            return levenshtein.Dist( here, goal )
    }
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
