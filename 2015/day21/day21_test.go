package day21

import(
    "testing"
    "github.com/stvp/assert"
    "github.com/schwern/adventofcode.go/combination"
    "github.com/schwern/adventofcode.go/util"
    "math"
)

func populateShop() *Shop {
    var weaponTexts = []string{
        `Dagger        8     4       0`,
        `Shortsword   10     5       0`,
        `Warhammer    25     6       0`,
        `Longsword    40     7       0`,
        `Greataxe     74     8       0`,
    }

    var armorTexts = []string{
        `Leather      13     0       1`,
        `Chainmail    31     0       2`,
        `Splintmail   53     0       3`,
        `Bandedmail   75     0       4`,
        `Platemail   102     0       5`,
        `NoArmor       0     0       0`,
    }

    var ringTexts = []string{
        `Damage+1    25     1       0`,
        `Damage+2    50     2       0`,
        `Damage+3   100     3       0`,
        `Defense+1   20     0       1`,
        `Defense+2   40     0       2`,
        `Defense+3   80     0       3`,
        `NoRing1      0     0       0`,
        `NoRing2      0     0       0`,
    }

    shop := NewShop()

    for _,weapon := range weaponTexts {
        shop.AddItem( ParseItem(weapon, Weapon) )
    }
    
    for _,armor := range armorTexts {
        shop.AddItem( ParseItem(armor, Armor) )
    }
    
    for _,ring := range ringTexts {
        shop.AddItem( ParseItem(ring, Ring) )
    }
    
    return shop
}

func TestParseItem( t *testing.T ) {
    item := ParseItem(`Damage+3   100     3       1`, Ring)
    
    assert.Equal( t, item.Name, "Damage+3" )
    assert.Equal( t, item.Cost, 100 )
    assert.Equal( t, item.Damage, 3 )
    assert.Equal( t, item.Armor, 1 )
    assert.Equal( t, item.Type, Ring )
}

func equipmentCombosChan( shop Shop ) chan []Item {
    ch := make( chan []Item )
    
    go func() {
        defer close(ch)
        
        weapons := shop.GetShelf( Weapon )
        armor   := shop.GetShelf( Armor )
        rings   := shop.GetShelf( Ring )
        
        for _,weapon := range weapons {
            for _,armor := range armor {
                for ringCombo := range combination.Chan( len(rings), 2 ) {                    
                    items := []Item{
                        weapon,
                        armor,
                        rings[ringCombo[0]],
                        rings[ringCombo[1]],
                    }

                    ch <- items
                }
            }
        }
    }()
    
    return ch
}

var boss = Stats{ HP: 109, Damage: 8, Armor: 2 }

func TestPart1( t *testing.T ) {
    shop := populateShop()

    bestCost := math.MaxInt32
    for combo := range equipmentCombosChan( *shop ) {
        player := NewPlayer(100, []int{1,1,2})
        cost := 0
        for _,item := range combo {
            player.EquipItem( item )
            cost += item.Cost
        }
                    
        if player.Fight(boss) {
            bestCost = util.MinInt( bestCost, cost )
        }
    }
    
    assert.Equal( t, bestCost, 111 )
}

func TestPart2( t *testing.T ) {
    shop := populateShop()

    worstCost := -1
    for combo := range equipmentCombosChan( *shop ) {
        player := NewPlayer(100, []int{1,1,2})
        cost := 0
        for _,item := range combo {
            player.EquipItem( item )
            cost += item.Cost
        }
                    
        if !player.Fight(boss) {
            worstCost = util.MaxInt( worstCost, cost )
        }
    }
    
    assert.Equal( t, worstCost, 188 )
}