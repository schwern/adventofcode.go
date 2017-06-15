package day22

import(
    _ "fmt"
    "errors"
    "math"
)

func startNextTurn( player *Player, boss *Boss ) {
    for name,effect := range player.Effects {
        boss.ApplySpellEffect(*effect)
        player.ApplySpellEffect(*effect)
        
        effect.Turns--
        if effect.Turns <= 0 {
            delete(player.Effects, name)
            player.RemoveSpellEffect(*effect)
        }
    }
}

func DoRound( player *Player, boss *Boss, spell Spell ) error {
    if player.HardMode {
        player.HP--
    }
    if player.HP <= 0 {
        return errors.New(`Player died`)
    }
    
    startNextTurn( player, boss )
    err := player.Cast(spell, boss)
    if err != nil {
        return err
    }
    if boss.HP <= 0 {
        return nil
    }
    
    startNextTurn( player, boss )
    if boss.HP > 0 {
        boss.Attack(player)
    }
    
    return nil
}

func incrementPath( path []int, idx int ) bool {
    if idx < 0 {
        return false
    }
    
    if path[idx] == len(Spells) - 1 {
        // Carry
        path[idx] = 0
        return incrementPath( path, idx-1 )
    } else {
        path[idx]++
    }
    
    return true
}

// Sets the path so the next call to incrementPath()
// will prune the path.
func prunePath( path []int, prunedAt int ) {
    for i := prunedAt+1; i < len(path); i++ {
        path[i] = len(Spells)-1
    }
}

// Brute force with pruning.
func LeastMana( player Player, boss Boss, maxTurns int ) int {
    least := math.MaxInt32

    path := make( []int, maxTurns )
    for {
        cost,prunedAt,err := tryPath( player, boss, path, least )
        // fmt.Printf("%v: %v,%v,%v\n", path, cost, prunedAt,err)
        if err == nil {
            if cost < least {
                least = cost
                prunePath(path, prunedAt)
            }
        } else {
            prunePath(path, prunedAt)
        }
        
        if !incrementPath(path, len(path)-1) {
            break
        }
    }
    
    return least
}

func tryPath(
    player Player, boss Boss, path []int, stop int,
) (int, int, error) {
    player.Reset()

    cost := 0
    for i,spellIdx := range path {
        spell := Spells[spellIdx]
        err := DoRound( &player, &boss, spell )
        cost += spell.Cost
        switch {
            case err != nil:
                return cost, i, err
            case cost >= stop:
                return cost, i, errors.New("Not cheaper")
            case player.HP <= 0:
                return cost, i, errors.New("Player died")
            case boss.HP <= 0:
                return cost,i,nil
        }
    }
    
    return cost, len(path), errors.New(`Boss still alive`)
}
