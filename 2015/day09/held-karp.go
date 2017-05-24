package day09

type HeldKarp struct {
    *Routes
    shortest [][]int
}

func NewHeldKarp( r *Routes ) HeldKarp {
    if r == nil {
        r = NewRoutes(false)
    }
    
    self := HeldKarp{
        Routes: r,
        shortest: r.edges,
    }
    
    return self
}

func (self *HeldKarp) ShortestRouteAny() (string, string, int) {
    return "", "", 0
}
