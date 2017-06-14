package day22

import(
    "errors"
)

type Player struct {
    HP int
    Mana int
    Armor int
    Effects map[string]SpellEffect
}

func NewPlayer( hp, mana, armor int ) *Player {
    effects := make( map[string]SpellEffect )
    self := Player{
        HP: hp,
        Mana: mana,
        Armor: armor,
        Effects: effects,
    }
    
    return &self
}

func (self *Player) Cast( spell Spell, boss *Boss ) error {
    if self.Mana < spell.Cost {
        return errors.New("Spell costs too much mana.")
    }
    
    instant := spell.Instant
    if instant != nil {
        boss.ApplySpellEffect(*spell.Instant)
        self.ApplySpellEffect(*spell.Instant)
    }
    
    overTime := spell.OverTime
    if overTime != nil {
        if _,ok := self.Effects[spell.Name]; ok {
            return errors.New("Spell already in effect.")
        }
        
        // Use a duplicate so decrementing the turns counter
        // won't change the main spell.
        self.Effects[spell.Name] = *overTime
    }
    
    return nil
}

func (self *Player) ApplySpellEffect( effect SpellEffect ) {
    self.Armor += effect.Armor
    self.HP += effect.Heal
    self.Mana += effect.Mana
}

func (self *Player) RemoveSpellEffect( effect SpellEffect ) {
    self.Armor -= effect.Armor
}
