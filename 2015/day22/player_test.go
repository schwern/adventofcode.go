package day22_test

import(
    "testing"
    "github.com/stvp/assert"
    "github.com/schwern/adventofcode.go/2015/day22"
)

func TestPossibleSpells( t *testing.T ) {
    spellBook := day22.GetSpellBook()
    emptySpellBook := make( day22.SpellBook )
    player := day22.NewPlayer( 10, 9999, 0, spellBook )
    boss := &day22.Boss{ HP: 999, Damage: 0 }
    
    player.Mana = 9999
    assert.Equal(
        t, player.PossibleSpells(), player.SpellBook,
        `New player with plenty of mana should be able to cast anything`,
    )
    
    player.Mana = 0
    assert.Equal(
        t, player.PossibleSpells(), emptySpellBook,
        `No mana, no spells should be able to cast anything`,
    )
    
    spellBookNoPoison := player.SpellBook
    delete( spellBookNoPoison, "Poison" )
    player.Mana = 9999
    player.Cast(player.SpellBook["Poison"], boss)
    assert.Equal(
        t, player.PossibleSpells(), spellBookNoPoison,
        `Spell already in effect unavailable`,
    )
}
