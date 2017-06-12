package day21

type ItemType int
const(
    Weapon ItemType = iota
    Armor
    Ring
)

type Item struct {
    Name string
    Cost int
    Damage int
    Armor int
    Type ItemType
}
