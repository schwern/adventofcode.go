package day15

import(
    "regexp"
    "github.com/schwern/adventofcode.go/util"
)

type Ingredient struct {
    Capacity int
    Durability int
    Flavor int
    Texture int
    Calories int
}

var ingredientRe = regexp.MustCompile(
    `\w+: capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)`,
)
func ParseIngredient( line string ) Ingredient {
    match := ingredientRe.FindStringSubmatch(line)
    if match == nil {
        util.Panicf("Cannot understand: %v", line)
    }
    
    return Ingredient{
        Capacity: util.MustAtoi(match[1]),
        Durability: util.MustAtoi(match[2]),
        Flavor: util.MustAtoi(match[3]),
        Texture: util.MustAtoi(match[4]),
        Calories: util.MustAtoi(match[5]),
    }
}

type Cookie struct {
    ingredients []Ingredient
}

func NewCookie() *Cookie {
    self := new(Cookie)
    self.ingredients = make( []Ingredient, 0 )
    return self
}

func (self *Cookie) AddIngredient( ing Ingredient ) {
    self.ingredients = append( self.ingredients, ing )
}

func (self *Cookie) Score( amounts []int ) int {
    cap := 0
    dur := 0
    fla := 0
    tex := 0
    for i, ing := range self.ingredients {
        cap += ing.Capacity * amounts[i]
        dur += ing.Durability * amounts[i]
        fla += ing.Flavor * amounts[i]
        tex += ing.Texture * amounts[i]
    }
    
    if cap < 0 || dur < 0 || fla < 0 || tex < 0 {
        return 0
    } else {
        return cap * dur * fla * tex
    }
}
