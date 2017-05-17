package day01

import "fmt"

func FindFloor(instructions string) (floor int) {
    floor = 0
    
    for _, char := range instructions {
        switch char {
            case '(':
                floor++
            case ')':
                floor--
            default:
                fmt.Printf("Unknown character '%v'\n", char)
        }
    }
    
    return
}

func FirstBasement(instructions string) (int) {
    floor := 0
    
    for pos, char := range instructions {
        switch char {
            case '(':
                floor++
            case ')':
                floor--
            default:
                fmt.Printf("Unknown character '%v'\n", char)
        }
        
        if floor < 0 {
            return pos + 1
        }
    }
    
    return 0
}
