package day21_test

import(
    "testing"
    "github.com/stvp/assert"
    "github.com/schwern/adventofcode.go/2015/day21"
)

func TestShop( t *testing.T ) {
    shop := day21.NewShop()
    emptyShelf := make( day21.ShopShelf )
    assert.Equal( t, shop.ListItems( day21.Weapon ), emptyShelf )
    
    spoon := day21.Item{
        Name: "Spoon!",
        Cost: 5,
        Damage: 10,
        Armor: 20,
        Type: day21.Weapon,
    }
    fork := day21.Item{
        Name: "Fork",
        Cost: 1,
        Damage: 2,
        Armor: 3,
        Type: day21.Weapon,
    }
    oneRing := day21.Item{
        Name: "The Ring",
        Cost: 10,
        Damage: 20,
        Armor: 30,
        Type: day21.Ring,
    }        
    
    shop.AddItem(spoon)
    shop.AddItem(fork)
    shop.AddItem(oneRing)
    
    wantWeapons := make( day21.ShopShelf )
    wantWeapons[spoon.Name] = spoon
    wantWeapons[fork.Name] = fork
    
    wantRings := make( day21.ShopShelf )
    wantRings[oneRing.Name] = oneRing
        
    assert.Equal( t, shop.ListItems( day21.Weapon ), wantWeapons )
    assert.Equal( t, shop.ListItems( day21.Ring ), wantRings )
    assert.Equal( t, shop.ListItems( day21.Armor ), emptyShelf )
    
    boughtSpoon := shop.BuyItem(spoon.Type, spoon.Name)
    delete( wantWeapons, spoon.Name )    
    assert.Equal( t, boughtSpoon, spoon )
    assert.Equal( t, shop.ListItems( day21.Weapon ), wantWeapons )
}
