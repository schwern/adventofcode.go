package day03

import "fmt"

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
            panic(fmt.Sprintf("Unknown direction %v", direction))
    }
}

func move( pos *[2]int, direction rune ) {
    dx, dy := parseDirection(direction)
    pos[0] += dx
    pos[1] += dy
}

func RoboPresentsDelivered( directions string ) int {
    count := 0
    grid := make( map[[2]int]int )
    x := 0
    y := 0
    
    santa := [2]int{x,y}
    robot := [2]int{x,y}
    grid[santa]++
    count++

    pos := &santa
    for i, dir := range directions {
        _ = "breakpoint"
        switch i % 2 {
            case 0: pos = &santa
            case 1: pos = &robot
            default: panic("What?")
        }
        move( pos, dir )
        if grid[*pos] == 0 {
            count++
        }
        grid[*pos]++
    }
    
    return count
}

func PresentsDelivered( directions string ) int {
    count := 0
    grid := make( map[[2]int]int )
    x := 0
    y := 0
    
    pos := [2]int{x,y}
    grid[pos]++
    count++
    
    for _, dir := range directions {        
        move( &pos, dir )        
        if grid[pos] == 0 {
            count++
        }
        grid[pos]++
    }
    
    return count
}
