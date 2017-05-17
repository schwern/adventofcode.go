package day02

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
