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

func ( r *Reindeer ) RunRunReindeer( time int ) (dist int) {
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

var reindeerRe = regexp.MustCompile(
    `\w+ can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`,
)
func ParseLine( line string ) *Reindeer {
    match := reindeerRe.FindStringSubmatch( line )
    if match == nil {
        util.Panicf("Can't understand: %s", line)
    }
    
    return &Reindeer{
        Speed: util.MustAtoi(match[1]),
        Duration: util.MustAtoi(match[2]),
        Rest: util.MustAtoi(match[3]),
    }
}
