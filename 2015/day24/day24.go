package day24

import(
    "math"
    "github.com/schwern/adventofcode.go/combination"
    "github.com/schwern/adventofcode.go/util"
)

func sumInts( nums []int ) (total int) {
    for _,num := range nums {
        total += num
    }
    
    return
}

func productInts( nums []int ) int {
    total := 1
    for _,num := range nums {
        total *= num
    }
    
    return total
}

func pickCombo( nums []int, combo []int ) []int {
    picks := make( []int, len(combo) )
    for pIdx,nIdx := range combo {
        picks[pIdx] = nums[nIdx]
    }
    return picks
}

func tryCombos( nums []int, k int, want int ) [][]int {
    validCombos := [][]int{}
    
    for combo := range combination.Chan(len(nums), k) {
        total := 0
        for _,i := range combo {
            total += nums[i]
        }
        if total == want {
            validCombos = append(
                validCombos,
                pickCombo(nums,combo),
            )
        }
    }
    
    return validCombos
}

func FindSmallestCombos( nums []int, N int ) [][]int {
    total := sumInts(nums)
    if total % N != 0 {
        panic("Can't divide evenly!")
    }
    
    want := total/N
    
    for k := 1; k < len(nums); k++ {
        combos := tryCombos( nums, k, want )
        if len(combos) != 0 {
            return combos
        }
    }
    
    return nil
}

func SmallestQE( combos [][]int ) int {
    qe := math.MaxInt64
    for _,combo := range combos {
        qe = util.MinInt( qe, productInts(combo) )
    }
    return qe
}
