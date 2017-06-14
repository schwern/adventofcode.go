package day22

import(
    "github.com/schwern/adventofcode.go/util"
)

type Boss struct {
    HP int
    Damage int
}

func (self *Boss) ApplySpellEffect( effect SpellEffect ) {
    self.HP -= effect.Damage
}

func (self *Boss) Attack( player *Player ) {
    dmg := util.MaxInt( self.Damage - player.Armor, 1 )
    player.HP -= dmg
}
