package day12

import(
    "encoding/json"
    "sync"
    "github.com/schwern/adventofcode.go/util"
)

func SumNums( in string, skip *string ) float64 {
    var data interface{}
    err := json.Unmarshal( []byte(in), &data )
    util.Check(err)
    
    out := make( chan float64 )    
    go func() {
        var wg sync.WaitGroup
        defer close(out)
        defer wg.Wait()
        
        wg.Add(1)
        go readNums( data, skip, out, &wg )
    }()
    
    total := 0.0
    for val := range out {
        total += val
    }
    
    return total
}

func readNums( data interface{}, skip *string, out chan float64, wg *sync.WaitGroup ) {
    defer wg.Done()
    
    switch d := data.(type) {
        case int64, float64:
            out <- data.(float64)
        case string:
            // ignore it
        case []interface{}:
            for _,u := range d {
                wg.Add(1)
                go readNums( u, skip, out, wg )
            }
        case map[string]interface{}:            
            if skip != nil {
                for _,u := range d {
                    if u == *skip {
                        return
                    }
                }
            }
            
            for _,u := range d {
                wg.Add(1)
                go readNums( u, skip, out, wg )
            }
        default:
            util.Panicf( "%v %T is of a type I don't know how to handle", data, data )
    }
}
