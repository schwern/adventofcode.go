package day14

import(
    "regexp"
    "github.com/schwern/adventofcode.go/util"
)

type Reindeer struct {
    Speed int
    Duration int
    Rest int
}

func ( r Reindeer ) RunRunReindeer( time int ) (dist int) {
    for time > r.Duration {
        dist += r.Speed * r.Duration
        time -= r.Duration
        time -= r.Rest
    }
    
    if time > 0 {
        dist += r.Speed * time
    }
    
    return
}

func ( r Reindeer ) isRunning( time int ) bool {
    if time < 0 {
        return false
    }
    return (time % (r.Duration + r.Rest )) < r.Duration
}

var reindeerRe = regexp.MustCompile(
    `\w+ can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`,
)
func ParseLine( line string ) Reindeer {
    match := reindeerRe.FindStringSubmatch( line )
    if match == nil {
        util.Panicf("Can't understand: %s", line)
    }
    
    return Reindeer{
        Speed: util.MustAtoi(match[1]),
        Duration: util.MustAtoi(match[2]),
        Rest: util.MustAtoi(match[3]),
    }
}

func ReindeerRace( rs []Reindeer, time int ) []int {
    points := make( []int, len(rs) )
    dists  := make( []int, len(rs) )
    
    for t := 0; t < time; t++ {
        for i,r := range rs {
            if r.isRunning(t) {
                dists[i] += r.Speed
            }
        }

        lead := util.MaxInts( dists )
        for i,dist := range dists {
            if dist == lead {
                points[i]++
            }
        }
    }
    
    return points
}
