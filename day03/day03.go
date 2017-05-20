package day03

import "fmt"

func move( direction rune ) (int, int) {
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

func PresentsDelivered( directions string ) int {
    count := 0
    grid := make( map[[2]int]int )
    x := 0
    y := 0
    
    pos := [2]int{x,y}
    grid[pos]++
    count++
    
    for _, dir := range directions {        
        dx, dy := move(dir)
        x += dx
        y += dy
        
        pos = [2]int{x,y}
        if grid[pos] == 0 {
            count++
        }
        grid[pos]++
    }
    
    return count
}
