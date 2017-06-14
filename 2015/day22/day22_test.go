package day22

import(
    "testing"
    "github.com/stvp/assert"
)

func TestExample1( t *testing.T ) {
    player := NewPlayer( 10, 250, 0 )
    boss := &Boss{ HP: 13, Damage: 8 }
    
    // The spell to cast, and what the player and boss
    // HP will be at the end of the boss' turn.
    tests := []struct{ spell Spell; playerHP, bossHP int }{
        { spell: SpellBook["Poison"], playerHP: 2, bossHP: 10 },
        { spell: SpellBook["Magic Missile"], playerHP: 2, bossHP: 0 },
    }
    
    for _,test := range tests {
        PlayerCast( player, boss, test.spell )
        assert.Equal( t, player.HP, test.playerHP )
        assert.Equal( t, boss.HP, test.bossHP )
    }
}

func TestExample2( t *testing.T ) {
    player := NewPlayer( 10, 250, 0 )
    boss := &Boss{ HP: 14, Damage: 8 }
    
    // The spell to cast, and what the player and boss
    // HP will be at the end of the boss' turn.
    tests := []struct{ spell Spell; playerHP, bossHP int }{
        { spell: SpellBook["Recharge"], playerHP: 2, bossHP: 14 },
        { spell: SpellBook["Shield"], playerHP: 1, bossHP: 14 },
        { spell: SpellBook["Drain"], playerHP: 2, bossHP: 12 },
        { spell: SpellBook["Poison"], playerHP: 1, bossHP: 9 },
        { spell: SpellBook["Magic Missile"], playerHP: 1, bossHP: -1 },
    }
    
    for _,test := range tests {
        PlayerCast( player, boss, test.spell )
        assert.Equal( t, player.HP, test.playerHP )
        assert.Equal( t, boss.HP, test.bossHP )
    }
}
