package day20

import(
    "runtime"
    "sync"
)

func DeliverPresents( maxHouse, maxDeliveries, numPresents int ) []int {
    presents := make( []int, maxHouse+1 )

    var wg sync.WaitGroup
    var mapMutex = &sync.Mutex{}
    deliver := func(start, by int) {
        defer wg.Done()
        
        for elf := start; elf <= maxHouse; elf+=by {
            numDelivered := 0
            
            // Putting the lock inside the loop is 10x slower.
            mapMutex.Lock()
            for houseNum := elf; houseNum <= maxHouse && numDelivered < maxDeliveries; houseNum += elf {
                presents[houseNum] += elf * numPresents
                numDelivered++
            }
            mapMutex.Unlock()
        }
    }

    startAt := 1
    numWorkers := runtime.NumCPU() + 1
    for i := startAt; i < numWorkers+startAt; i++ {
        wg.Add(1)
        go deliver(i, numWorkers)
    }
    wg.Wait()
    
    return presents
}
