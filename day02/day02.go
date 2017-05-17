package day02

import "sort"

func min( nums ...int ) (min int) {
    min = int(^uint(0) >> 1)
    for _, i := range nums {
        if i < min {
            min = i
        }
    }
    
    return min
}

func WrappingPaperNeeded( l, w, h int ) (paper int) {
    sides := []int{ l*w, w*h, h*l }
    paper = min(sides...)
    
    for _, side := range sides {
        paper += 2*side
    }
    
    return
}

func RibbonNeeded( l, w, h int ) (ribbon int) {
    sides := []int{l, w, h}
    sort.Ints( sides )
    ribbon = 2*sides[0] + 2*sides[1] + l * w * h
    return
}