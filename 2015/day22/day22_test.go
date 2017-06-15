package day22

import(
    "testing"
    "github.com/stvp/assert"
)

func TestExample1( t *testing.T ) {
    spellBook := GetSpellBook()
    player := NewPlayer( 10, 250, 0, spellBook )
    boss := &Boss{ HP: 13, Damage: 8 }
    
    // The spell to cast, and what the player and boss
    // HP will be at the end of the boss' turn.
    tests := []struct{ spell Spell; playerHP, bossHP, mana int }{
        { spell: spellBook["Poison"], playerHP: 2, bossHP: 10, mana: 77 },
        { spell: spellBook["Magic Missile"], playerHP: 2, bossHP: 0, mana: 24 },
    }
    
    for _,test := range tests {
        DoRound( player, boss, test.spell )
        assert.Equal( t, player.HP, test.playerHP )
        assert.Equal( t, player.Mana, test.mana )
        assert.Equal( t, boss.HP, test.bossHP )
    }
}

func TestTryPath( t *testing.T ) {
    sb := GetSpellBook()
    player := NewPlayer( 10, 250, 0, sb )
    boss := &Boss{ HP: 13, Damage: 8 }
        
    path := []int{ 3, 0, 0 }
    cost, prunedAt, err := tryPath( *player, *boss, path, 266 )
    assert.Nil( t, err )
    assert.Equal( t, cost, 226 )
    assert.Equal( t, prunedAt, 1 )
}

func TestLeastMana( t *testing.T ) {
    spellBook := GetSpellBook()
    player := NewPlayer( 10, 250, 0, spellBook )
    boss := &Boss{ HP: 13, Damage: 8 }
    
    assert.Equal( t, LeastMana(*player, *boss, 5), 226 )
}

func TestIncrementPath( t *testing.T ) {
    maxIdx := len(Spells) - 1
    path := []int{1,1,1}
    assert.True( t, incrementPath(path, 0) )
    assert.Equal( t, path, []int{2,1,1}, "Increment" )
    
    path = []int{0,maxIdx,maxIdx,maxIdx}
    assert.True( t, incrementPath(path, 3) )
    assert.Equal( t, path, []int{1,0,0,0}, "Carry" )
        
    path = []int{maxIdx,maxIdx,maxIdx}
    assert.False( t, incrementPath(path,2), "End" )    
}

func TestPrunePath( t *testing.T ) {
    maxIdx := len(Spells) - 1
    path := []int{1,2,3,1,2}
    prunePath( path, 2 )
    
    assert.Equal( t, path, []int{1,2,3,maxIdx,maxIdx} )
    incrementPath( path, len(path)-1 )
    assert.Equal( t, path, []int{1,2,4,0,0} )
}

func TestExample2( t *testing.T ) {
    spellBook := GetSpellBook()
    player := NewPlayer( 10, 250, 0, spellBook )
    boss := &Boss{ HP: 14, Damage: 8 }
    
    // The spell to cast, and what the player and boss
    // HP will be at the end of the boss' turn.
    tests := []struct{ spell Spell; playerHP, bossHP int }{
        { spell: spellBook["Recharge"], playerHP: 2, bossHP: 14 },
        { spell: spellBook["Shield"], playerHP: 1, bossHP: 14 },
        { spell: spellBook["Drain"], playerHP: 2, bossHP: 12 },
        { spell: spellBook["Poison"], playerHP: 1, bossHP: 9 },
        { spell: spellBook["Magic Missile"], playerHP: 1, bossHP: -1 },
    }
    
    for _,test := range tests {
        DoRound( player, boss, test.spell )
        assert.Equal( t, player.HP, test.playerHP )
        assert.Equal( t, boss.HP, test.bossHP )
    }
}

func TestPart1( t *testing.T ) {
    spellBook := GetSpellBook()
    player := NewPlayer( 50, 500, 0, spellBook )
    boss := &Boss{ HP: 55, Damage: 8 }
    
    assert.Equal( t, LeastMana( *player, *boss, 10 ), 953 )
}