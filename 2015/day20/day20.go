package day20

func DeliverPresents( maxHouse, maxDeliveries, numPresents int ) []int {
    presents := make( []int, maxHouse+1 )

    for elf := 1; elf <= maxHouse; elf++ {
        numDelivered := 0
            
        for houseNum := elf; houseNum <= maxHouse && numDelivered < maxDeliveries; houseNum += elf {
            presents[houseNum] += elf * numPresents
            numDelivered++
        }
    }

    return presents
}
