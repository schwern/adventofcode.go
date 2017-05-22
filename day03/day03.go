package day03

import(
    "github.com/schwern/adventofcode2015/util"
)

func parseDirection( direction rune ) (int, int) {
    switch direction {
        case '^':
            return 0, 1
        case 'v':
            return 0, -1
        case '<':
            return -1, 0
        case '>':
            return 1, 0
        default:
            util.Panicf("Unknown direction %v", direction)
            return 0,0
    }
}

func move( pos *[2]int, direction rune ) {
    dx, dy := parseDirection(direction)
    pos[0] += dx
    pos[1] += dy
}

func DeliverPresents( directions string, num_movers int ) int {
    count := 0
    grid := make( map[[2]int]int )
    movers := make( [][2]int, num_movers )
    
    grid[movers[0]]++
    count++

    var pos *[2]int
    for i, dir := range directions {
        pos = &movers[i % num_movers]
        move( pos, dir )
        if grid[*pos] == 0 {
            count++
        }
        grid[*pos]++
    }
    
    return count
}
