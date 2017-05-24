package day09

import(
    "github.com/schwern/adventofcode.go/util"
)

type RouteKey [2]string
type Routes map[RouteKey]int

func (self Routes) SetRoute( a, b string, dist int ) {
    self[ makeKey( a, b ) ] = dist
}

func (self Routes) GetRoute( a, b string ) int {
    return self[ makeKey( a, b ) ]
}

func NewRoutes() Routes {
    return make( Routes )
}

func makeKey( a,b string ) RouteKey {
    switch {
        case a < b:
            return [2]string{ a, b }
        case b < a:
            return [2]string{ b, a }
        default: 
            util.Panicf("keys are equal: %v", a)
            return [2]string{}
    }
}
