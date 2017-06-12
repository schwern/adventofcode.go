package day21_test

import(
    "testing"
    "github.com/stvp/assert"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/2015/day21"
)

func TestEquipItem( t *testing.T ) {
    limits := make( []int, 3 )
    limits[day21.Weapon] = 1
    limits[day21.Armor] = 1
    limits[day21.Ring] = 2
        
    player := day21.NewPlayer( 100, limits )
    
    spoon := day21.Item{
        Name: "Spoon!",
        Cost: 5,
        Damage: 10,
        Armor: 20,
        Type: day21.Weapon,
    }
    
    babies := day21.Item{
        Name: "Suit Of Babies",
        Cost: 100,
        Damage: 0,
        Armor: 50,
        Type: day21.Armor,
    }

    player.EquipItem( spoon )
    assert.Equal( t, player.Damage, spoon.Damage )
    assert.Equal( t, player.Armor, spoon.Armor )
    
    player.EquipItem( babies )
    assert.Equal( t, player.Damage, spoon.Damage + babies.Damage )
    assert.Equal( t, player.Armor, spoon.Armor + babies.Armor )
    
    defer testutil.AssertPanicf( t, "Can't equip more items of this type, already have 1" )
    player.EquipItem( spoon )
}
