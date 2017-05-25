package day10

import(
    "strconv"
)

func LookSay( in []byte ) []byte {
    out := make( []byte, 0, len(in) )
    prev := in[0]
    count := int64(1)
    for i := 1; i <= len(in); i++ {
        if i >= len(in) || in[i] != prev {
            out = strconv.AppendInt(out, count, 10)
            out = append( out, prev )

            count = 1
            if i < len(in) {
                prev = in[i]
            }
        } else {
            count++
        }
    }
    
    return out
}
