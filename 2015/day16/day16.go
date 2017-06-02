package day16

import(
    "regexp"
    "strings"
    "github.com/schwern/adventofcode.go/util"
)

type Compounds map[string]int
type Sue struct {
    Num int
    compounds Compounds
}

func NewSue( num int, compunds Compounds ) *Sue {
    sue := Sue{ Num: num, compounds: compunds }
    return &sue
}

var sueRe = regexp.MustCompile(
    `Sue (\d+): (.+)`,
)
func ParseSue( line string ) *Sue {
    match := sueRe.FindStringSubmatch( line )
    if match == nil {
        util.Panicf("Cannot understand: %v", line)
    }
    
    num := match[1]
    
    compounds := ParseCompunds( match[2] )
    
    sue := Sue{ Num: util.MustAtoi(num), compounds: compounds }
    return &sue
}

func ParseCompunds( line string ) Compounds {
    compounds := make( Compounds )
    for _,compound := range strings.Split( line, ", " ) {
        pair := strings.SplitN( compound, ": ", 2 )
        compounds[pair[0]] = util.MustAtoi(pair[1])
    }
    
    return compounds
}

func (self *Sue) CheckCompounds( want Compounds ) bool {
    compounds := self.compounds
    for key,have := range compounds {
        if have != want[key] {
            return false
        }
    }
    
    return true
}
