package day09

import(
    "math"
    "sync"
    "regexp"
    "github.com/schwern/adventofcode.go/util"
)

func BruteForce( routes *Routes ) int {
    nodes := make( []int, routes.NumNodes() )
    for i := range nodes {
        nodes[i] = i
    }
    
    perms := NewPermutationChan( nodes )
    ch := make( chan int )
    
    var wg sync.WaitGroup
    for perm := range perms {
        wg.Add(1)
        go func() {
            defer wg.Done()
            total := 0
            for i := 1; i < len(perm); i++ {
                dist := routes.GetRouteByIdx( i-1, i )
                if dist == 0 {
                    ch <- math.MaxInt32
                    return
                } else {
                    total += dist
                }
            }
            
            ch <- total
        }()
    }
    
    go func() {
        wg.Wait()
        close(ch)
    }()
    
    shortest := math.MaxInt32
    for dist := range ch {
        if dist < shortest {
            shortest = dist
        }
    }
    
    return shortest
}


var lineRe = regexp.MustCompile(
    `(\w+) to (\w+) = (\d+)`,
)
func ParseLine( line string ) (string, string, int) {
    match := lineRe.FindStringSubmatch(line)

    return match[1], match[2], util.MustAtoi(match[3])
}
