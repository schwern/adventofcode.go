package day17

type Containers []int

func (self Containers) NumCombos( want int ) int {
    numCombos := 0
    for _ = range self.validComboChan(want) {
        numCombos++
    }
    
    return numCombos
}

func (self Containers) NumMinCombos( want int ) int {
    comboSize := len(self)+1
    combos := 0
    
    for combo := range self.validComboChan(want) {
        size := self.numContainersInCombo(combo)
        switch {
            case size > comboSize:
                // do nothing
            case size < comboSize:
                comboSize = size
                combos = 1
            case size == comboSize:
                combos++
        }
    }
    
    return combos
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

func isInCombo( combo uint, i int ) bool {
    return combo & (1<<uint(i)) != 0
}

func (self Containers) numContainersInCombo( combo uint ) int {
    size := 0
    for i := range self {
        if isInCombo(combo, i) {
            size++
        }
    }
    
    return size
}

func (self Containers) combo( combo uint ) int {
    storage := 0
    for i := range self {
        if isInCombo(combo, i) {
            storage += self[i]
        }
    }
    
    return storage
}
