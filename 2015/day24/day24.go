package day24

import(
    "github.com/schwern/adventofcode.go/combination"
)

func sumInts( nums []int ) (total int) {
    for _,num := range nums {
        total += num
    }
    
    return
}

func smallestBucketsChan( nums []int ) chan []int {
    total := sumInts(nums)
    if total % 3 != 0 {
        panic("Can't divide evenly!")
    }
    
    want := total/3
    
    ch := make( chan []int )
    go func() {
        defer close(ch)
        
        for k := 1; k < len(nums); k++ {
            combos := tryCombos( nums, k, want )
            for _,combo := range combos {
                ch <- combo
            }
        }
    }()
    
    return ch
}

func tryCombos( nums []int, k int, want int ) [][]int {
    validCombos := [][]int{}
    
    for combo := range combination.Chan(len(nums), k) {
        total := 0
        for _,i := range combo {
            total += nums[i]
        }
        if total == want {
            validCombos = append(validCombos, combo)
        }
    }
    
    return validCombos
}
