package day10

import(
    "strconv"
)

func LookSay( in string ) string {
    out := ""
    prev := in[0]
    count := 1
    for i := 1; i <= len(in); i++ {
        if i >= len(in) || in[i] != prev {
            out += strconv.Itoa(count)
            out += string(prev)

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
