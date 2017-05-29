package routes

import(
    "sync"
    "github.com/schwern/adventofcode.go/permutation"
    "github.com/schwern/adventofcode.go/util"
)

type errorNoRoute struct {}
func (self *errorNoRoute) Error() string {
    return "No route exists"
}

type errorNoPath struct {}
func (self *errorNoPath) Error() string {
    return "No path exists"
}

type Routes struct {
    edges [][]*int
    names map[string]int
    is_symetric bool
}

func NewRoutes( symetric bool ) *Routes {
    self := Routes{
        edges: make( [][]*int, 0 ),
        names: make( map[string]int ),
    }
    
    self.is_symetric = symetric
    
    return &self
}

func (self *Routes) addNode( name string ) int {
    len := len(self.edges)
    
    // Extend each existing row by 1    
    for i := range self.edges {
        self.edges[i] = append(self.edges[i], nil)
    }
    
    // Add a new row
    self.edges = append(self.edges, make( []*int, len+1) )
    
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

func (self *Routes) AddRoute( a, b string, cost int ) {
    aidx := self.getOrAddNodeIdx( a )
    bidx := self.getOrAddNodeIdx( b )
    
    self.edges[ aidx ][ bidx ] = &cost
    if self.is_symetric {
        self.edges[ bidx ][ aidx ] = &cost
    }
}

func (self *Routes) GetRoute( a, b string ) (int, error) {
    return self.GetRouteByIdx( self.mustGetNodeIdx(a), self.mustGetNodeIdx(b) )
}

func (self *Routes) MustGetRoute( a, b string ) int {
    cost, err := self.GetRoute( a, b )
    if err != nil {
        panic(err)
    }
    return cost
}

func (self *Routes) GetRouteByIdx( a, b int ) (int, error) {
    val := self.edges[a][b]
    
    if val == nil {
        return 0, new(errorNoRoute)
    } else {
        return *val, nil
    }
}

func (self *Routes) NumNodes() int {
    return len(self.edges)
}

func (self *Routes) PathCost( path []int ) (int, error) {
    total := 0

    var error error
    for i := 1; i < len(path); i++ {
        a := path[i-1]
        b := path[i]
        cost, err := self.GetRouteByIdx( a, b )
                        
        if err != nil {
            error = new(errorNoPath)
            total = 0
            break
        } else {
            total += cost
        }
    }
        
    return total, error
}

func (self *Routes) AllPathsChan() chan []int {
    nodes := make( []int, self.NumNodes() )
    for i := range nodes {
        nodes[i] = i
    }
    return permutation.NewPermutationChan( nodes )
}

func (self *Routes) TryAllPaths() chan int {    
    return self.TryPaths( self.AllPathsChan() )
}

func (self *Routes) TryPaths( paths chan []int ) chan int {
    ch := make( chan int )
    
    var wg sync.WaitGroup
    for path := range paths {
        wg.Add(1)
        
        go func( path []int ) {
            defer wg.Done()
            cost, err := self.PathCost( path )
            if err != nil {
                return
            }
            ch <- cost
        }(path)
    }
    
    go func() {
        wg.Wait()
        close(ch)
    }()
    
    return ch
}
