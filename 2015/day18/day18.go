package day18

import(
    "github.com/schwern/adventofcode.go/util"
)

func parseLine( line string ) (row []bool) {
    for _,c := range line {
        var state bool
        switch c {
            case '.':
                state = false
            case '#':
                state = true
            default:
                util.Panicf("Unknown state: %v", state)
        }
        
        row = append( row, state )
    }
    
    return
}

func ParseGridChan( lines chan string ) (grid Grid) {
    for line := range lines {
        grid = append( grid, parseLine(line) )
    }
    
    return
}

func ParseGrid( lines []string ) (grid Grid) {
    for _,line := range lines {
        grid = append( grid, parseLine(line) )
    }
    
    return
}
