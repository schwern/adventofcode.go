package day13

import(
    "regexp"
    "github.com/schwern/adventofcode.go/util"
    _ "github.com/schwern/adventofcode.go/permutation"
    _ "github.com/schwern/adventofcode.go/routes"
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
