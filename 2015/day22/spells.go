package day22

type Spell struct {
    Name string
    Cost int
    Instant *SpellEffect
    OverTime *SpellEffect
}

type SpellEffect struct {
    Turns int
    Armor int
    Damage int
    Heal int
    Mana int
}

var Spells = []Spell{
    Spell{
        Name: "Magic Missile",
        Cost: 53,
        Instant: &SpellEffect{ Damage: 4 },
    },
    Spell{
        Name: "Drain",
        Cost: 73,
        Instant: &SpellEffect{ Damage: 2, Heal: 2 },
    },
    Spell{
        Name: "Shield",
        Cost: 113,
        OverTime: &SpellEffect{ Turns: 6, Armor: 7 },
    },
    Spell{
        Name: "Poison",
        Cost: 173,
        OverTime: &SpellEffect{ Turns: 6, Damage: 3 },
    },
    Spell{
        Name: "Recharge",
        Cost: 229,
        OverTime: &SpellEffect{ Turns: 5, Mana: 101 },
    },
}

var spellBook = makeSpellBook( Spells )
type SpellBook map[string]Spell

func GetSpellBook() SpellBook {
    return spellBook
}

func makeSpellBook( spells []Spell ) SpellBook {
    book := make( SpellBook )
    
    for _,spell := range spells {
        book[spell.Name] = spell
    }
    
    return book
}

func (self *SpellEffect) Duplicate() *SpellEffect {
    copy := &SpellEffect{}
    *copy = *self
    return copy
}
