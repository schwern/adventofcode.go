package main

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