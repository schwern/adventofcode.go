package day15

import(
    "math"
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

func (self *Cookie) BestScore( total int ) int {
    best := math.MinInt32
    combos := self.ingredientComboChan( total )
    for combo := range combos {
        best = util.MaxInt( best, self.Score(combo) )
    }
    
    return best
}

func (self *Cookie) ingredientComboChan( total int ) chan []int {    
    ch := make( chan []int )
    go func() {
        defer close(ch)
        
        combo := make( []int, len(self.ingredients) )
        combo[0] = total
        
        for combo != nil {
            next := make( []int, len(combo) )
            copy( next, combo )
            ch <- next
            combo = self.nextIngredientCombo( combo, total )
        }
    }()
    
    return ch
}

func (self *Cookie) nextIngredientCombo( combo []int, total int ) []int {
    sum := 0
    lastIdx := len(combo) - 1
    
    // We're done.
    if combo[lastIdx] == total {
        return nil
    }
    
    for i := lastIdx; i >= 0; i-- {
        sum += combo[i]
        
        if sum == total {
            // Send one up
            combo[i+1]++
            combo[i]--
            
            // Send the remainder down
            if combo[i] > 0 {            
                downIdx := util.MaxInt( 0, i-1 )
                remainder := combo[i]
                combo[i] = 0
                combo[downIdx] = remainder
            }
            
            break
        }
    }
    
    return combo
}
