package day21

import(
    "github.com/schwern/adventofcode.go/util"
)

type Stats struct {
    HP int
    Damage int
    Armor int
}

func (self Stats) Attack( target Stats ) int {
    // Always do at least 1 HP of damage
    return util.MaxInt( self.Damage - target.Armor, 1 )
}

type Player struct {
    Stats
    Slots []int
    MaxSlots []int
}

func NewPlayer( hp int, maxSlots []int ) *Player {
    self := Player{
        Stats: Stats{ HP: hp, Damage: 0, Armor: 0 },
        MaxSlots: maxSlots,
        Slots: make( []int, len(maxSlots) ),
    }
    
    return &self
}

func (self *Player) EquipItem( item Item ) {
    if self.Slots[item.Type] >= self.MaxSlots[item.Type] {
        util.Panicf(
            "Already have %v of type %v equipped",
            self.MaxSlots[item.Type],
            item.Type,
        )
    }
    
    self.Slots[item.Type]++
    
    self.Damage += item.Damage
    self.Armor  += item.Armor
}

func (self *Player) UnequipItem( item Item ) {
    if self.Slots[item.Type] <= 0 {
        util.Panicf(
            "No items of type %v equipped.", item.Type,
        )
    }
    
    self.Slots[item.Type]--
    self.Damage -= item.Damage
    self.Armor  -= item.Armor
}

// Fight an opponent TO THE DEATH!
// Returns whether the opponent won or not.
// Does not alter the player nor target.
func (self Player) Fight( target Stats ) bool {
    selfRatio := float64(self.HP) / float64(target.Attack(self.Stats))
    targetRatio := float64(target.HP) / float64(self.Attack(target))
    
    return selfRatio >= targetRatio
}
