package combination

func dup( a []int ) []int {
    b := make( []int, len(a) )
    copy( b, a )
    return b
}

func Chan( N, K int ) chan []int {
    c := make( []int, K )
    ch := make( chan []int )
    
    var combos func( int, int )
    combos = func( start, k int ) {
        if k >= K {
            ch <- dup(c)
            return
        }
        
        for i := start; i < N; i++ {
            c[k] = i
            
            combos( i+1, k+1 )
        }
    }
    
    go func() {
        defer close(ch)
        combos( 0, 0 )
    }()
    
    return ch
}
