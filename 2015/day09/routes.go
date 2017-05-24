package day09

import(
    "github.com/schwern/adventofcode.go/util"
)

type Routes struct {
    edges [][]int
    names map[string]int
    is_symetric bool
}

func NewRoutes( symetric bool ) Routes {
    self := Routes{
        edges: make( [][]int, 0 ),
        names: make( map[string]int ),
    }
    
    self.is_symetric = symetric
    
    return self
}

func (self *Routes) addNode( name string ) int {
    len := len(self.edges)
    
    // Extend each existing row by 1    
    for i := range self.edges {
        self.edges[i] = append(self.edges[i], 0)
    }
    
    // Add a new row
    self.edges = append(self.edges, make( []int, len+1) )
    
    self.names[name] = len
    
    return len
}

func (self *Routes) getNodeIdx( name string ) (int, bool) {
    idx, exists := self.names[name]
    return idx, exists
}

func (self *Routes) mustGetNodeIdx( name string ) (int) {
    idx, exists := self.getNodeIdx(name)
    if !exists {
        util.Panicf("Do not have a node named %v", name)
        return 0
    }
    
    return idx
}

func (self *Routes) getOrAddNodeIdx( name string ) int {
    idx, exists := self.getNodeIdx( name )
    if !exists {
        idx = self.addNode( name )
    }

    return idx
}

func (self *Routes) AddRoute( a, b string, dist int ) {
    aidx := self.getOrAddNodeIdx( a )
    bidx := self.getOrAddNodeIdx( b )
    
    self.edges[ aidx ][ bidx ] = dist
    if self.is_symetric {
        self.edges[ bidx ][ aidx ] = dist
    }
}

func (self *Routes) GetRoute( a, b string ) int {
    return self.edges[ self.mustGetNodeIdx(a) ][ self.mustGetNodeIdx(b) ]
}
