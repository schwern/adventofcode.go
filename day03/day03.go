package day03

import "fmt"

func PresentsDelivered( directions string ) int {
    count := 0
    grid := make( map[[2]int]int )
    x := 0
    y := 0
    
    pos := [2]int{x,y}
    grid[pos]++
    count++
    
    for _, dir := range directions {        
        switch dir {
            case '^':
                y++
            case 'v':
                y--
            case '<':
                x--
            case '>':
                x++
            default:
                panic(fmt.Sprintf("Unknown direction %v", dir))
        }
        
        pos = [2]int{x,y}
        if grid[pos] == 0 {
            count++
        }
        grid[pos]++
    }
    
    return count
}