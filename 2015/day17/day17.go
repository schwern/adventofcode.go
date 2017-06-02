package day17

type Containers []int

func (self Containers) NumCombos( want int ) int {
    numCombos := 0
    for _ = range self.validComboChan(want) {
        numCombos++
    }
    
    return numCombos
}

func (self Containers) validComboChan( want int ) chan uint {
    ch := make(chan uint)
    
    go func() {
        defer close(ch)
        
        var i uint
        for i = 0; i < (1 << uint(len(self))); i++ {
            if self.combo( i ) == want {
                ch <- i
            }
        }
    }()
    
    return ch
}

func (self Containers) combo( combo uint ) int {
    storage := 0
    for i := range self {
        if combo & (1<<uint(i)) != 0 {
            storage += self[i]
        }
    }
    
    return storage
}
