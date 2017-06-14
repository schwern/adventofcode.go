package day22

func startNextTurn( player *Player, boss *Boss ) {
    for name,effect := range player.Effects {
       boss.ApplySpellEffect(effect)
       player.ApplySpellEffect(effect)
        
        effect.Turns--
        if effect.Turns <= 0 {
            delete(player.Effects, name)
        }
    }
}

func PlayerCast( player *Player, boss *Boss, spell Spell ) {
    startNextTurn( player, boss )
    player.Cast(spell, boss)
    startNextTurn( player, boss )
    if boss.HP > 0 {
        boss.Attack(player)
    }
}
