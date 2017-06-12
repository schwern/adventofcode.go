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
    Equipment [][]Item
}

func NewPlayer( hp int, limits []int ) *Player {
    self := Player{
        Stats: Stats{ HP: hp, Damage: 0, Armor: 0 },
        Equipment: make( [][]Item, len(limits) ),
    }
    
    for i,limit := range limits {
        self.Equipment[i] = make( []Item, 0, limit )
    }
    
    return &self
}

func (self *Player) EquipItem( item Item ) {
    items := self.Equipment[item.Type]
    if cap(items) == len(items) {
        util.Panicf("Can't equip more items of this type, already have %v", len(items))
    }
    
    self.Equipment[item.Type] = append( items, item )
    self.Damage += item.Damage
    self.Armor  += item.Armor
}

// Fight an opponent TO THE DEATH!
// Returns whether the opponent won or not.
// Does not alter the player nor target.
func (self Player) Fight( target Stats ) bool {
    selfRatio := float64(self.HP) / float64(target.Attack(self.Stats))
    targetRatio := float64(target.HP) / float64(self.Attack(target))
    
    return selfRatio >= targetRatio
}
