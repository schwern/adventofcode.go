package day13

import(
    "math"
    "regexp"
    "github.com/schwern/adventofcode.go/util"
    "github.com/schwern/adventofcode.go/routes"
)

var lineRe = regexp.MustCompile(
    `(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+)\.`,
)

func ParseLine( line string ) ( a, b string, dist int ) {
    match := lineRe.FindStringSubmatch(line)
    if match == nil {
        util.Panicf("Can't understand: %v", line)
    }
    
    a = match[1]
    b = match[4]
    dist = util.MustAtoi(match[3])
    
    switch match[2] {
        case "gain":
            // nothing
        case "lose":
            dist *= -1
        default:
            util.Panicf("Unknown happiness change: %v", match[1])
    }
    
    return
}

func thereAndBack( list []int ) []int {
    new := make( []int, len(list)*2 + 1 )
    copy( new, list )
    
    // Complete the trip around the table.
    new[len(list)] = list[0]
    
    // Go back to the beginning
    lastIdx := len(new)-1
    for i,v := range list {
        new[lastIdx-i] = v
    }
    
    return new
}

func allSeatingChan( r *routes.Routes ) chan []int {
    paths := r.AllPathsChan()
    returnPaths := make( chan []int )
    go func() {
        defer close(returnPaths)
        
        for path := range paths {
            returnPaths <- thereAndBack(path)
        }
    }()
    
    return returnPaths
}

func HappiestSeating( r *routes.Routes ) int {
    happiest := math.MinInt32
    for happy := range r.TryPaths( allSeatingChan(r) ) {
        happiest = util.MaxInt(happiest, happy )
    }
    
    return happiest
}
