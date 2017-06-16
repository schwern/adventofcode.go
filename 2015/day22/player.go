package day22

import(
    "errors"
)

type Player struct {
    HP int
    Mana int
    Armor int
    HardMode bool
    Effects map[string]*SpellEffect
    SpellBook SpellBook
}

func NewPlayer( hp, mana, armor int, spellBook SpellBook ) *Player {
    self := Player{
        HP: hp,
        Mana: mana,
        Armor: armor,
        SpellBook: spellBook,
    }
    
    self.Reset()
    
    return &self
}

func (self *Player) Reset() {
    self.Effects = make( map[string]*SpellEffect )
}

func (self *Player) PossibleSpells() SpellBook {
    possible := make( SpellBook )
    
    for name,spell := range self.SpellBook {
        if spell.Cost > self.Mana {
            continue
        }
        if _,ok := self.Effects[name]; ok {
            continue
        }
        
        possible[name] = spell
    }
    
    return possible
}

func (self *Player) Cast( spell Spell, boss *Boss ) error {
    if self.Mana < spell.Cost {
        return errors.New("Spell costs too much mana.")
    }
    self.Mana -= spell.Cost
    
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
        dup := &SpellEffect{}
        *dup = *overTime
        self.Effects[spell.Name] = dup
    }
    
    return nil
}

func (self *Player) ApplySpellEffect( effect SpellEffect ) {
    if effect.Armor > 0 {
        self.Armor = effect.Armor
    }
    self.HP += effect.Heal
    self.Mana += effect.Mana
}

func (self *Player) RemoveSpellEffect( effect SpellEffect ) {
    self.Armor -= effect.Armor
}
