package day09

import(
    "math"
    "sync"
    "regexp"
    "github.com/schwern/adventofcode.go/permutation"
    "github.com/schwern/adventofcode.go/routes"
    "github.com/schwern/adventofcode.go/util"
)

func bruteForce( routes *routes.Routes ) chan int {
    nodes := make( []int, routes.NumNodes() )
    for i := range nodes {
        nodes[i] = i
    }
        
    perms := permutation.NewPermutationChan( nodes )
    ch := make( chan int )
    
    var wg sync.WaitGroup
    for perm := range perms {
        wg.Add(1)
        
        go func( perm []int ) {
            defer wg.Done()
            ch <- routes.PathCost( perm )
        }(perm)
    }
    
    go func() {
        wg.Wait()
        close(ch)
    }()
    
    return ch
}

func min( a, b int ) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func max( a, b int ) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func BestRouteBruteForce( routes *routes.Routes ) int {
    ch := bruteForce( routes )
    
    shortest := math.MaxInt32
    for dist := range ch {
        if dist > 0 {
            shortest = min(dist, shortest)
        }
    }
    
    return shortest
}

func WorstRouteBruteForce( routes *routes.Routes ) int {
    ch := bruteForce( routes )
    
    longest := 0
    for dist := range ch {
        if dist > 0 {
            longest = max(dist, longest)
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
