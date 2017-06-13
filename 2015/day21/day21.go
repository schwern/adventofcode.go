package day21

import(
    "regexp"
    "github.com/schwern/adventofcode.go/util"
)

var whitespaceRe = regexp.MustCompile(`\s+`)
func ParseItem( line string, itemType ItemType ) Item {
    fields := whitespaceRe.Split(line, -1)
    
    return Item{
        Name: fields[0],
        Cost: util.MustAtoi(fields[1]),
        Damage: util.MustAtoi(fields[2]),
        Armor: util.MustAtoi(fields[3]),
        Type: itemType,
    }
}
