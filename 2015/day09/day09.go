package day09

import(
    "math"
    "regexp"
    "github.com/schwern/adventofcode.go/routes"
    "github.com/schwern/adventofcode.go/util"
)

func BestRouteBruteForce( routes *routes.Routes ) int {
    ch := routes.TryAllPaths()
    
    shortest := math.MaxInt32
    for dist := range ch {
        if dist > 0 {
            shortest = util.MinInt(dist, shortest)
        }
    }
    
    return shortest
}

func WorstRouteBruteForce( routes *routes.Routes ) int {
    ch := routes.TryAllPaths()
    
    longest := 0
    for dist := range ch {
        if dist > 0 {
            longest = util.MaxInt(dist, longest)
        }
    }
    
    return longest
}

var lineRe = regexp.MustCompile(
    `(\w+) to (\w+) = (\d+)`,
)
func ParseLine( line string ) (string, string, int) {
    match := lineRe.FindStringSubmatch(line)

    return match[1], match[2], util.MustAtoi(match[3])
}
