package day08

import(
    "unicode"
    "github.com/schwern/adventofcode.go/util"
)

func EncodingSize( str string ) int {
    // Start with the surrounding quotes
    size := 2
    
    for _,c := range str {
        if unicode.IsDigit(c) || unicode.IsLetter(c) {
            size++
        } else {
            size += 2
        }
    }
    
    return size
}

func MemorySize( str string ) int {
    size := 0
    
    for i := 0; i < len(str); i++ {
        size++
        
        switch str[i] {
            case '\\':
                switch c := str[i+1]; c {
                    // \\ or \"
                    case '\\', '"':
                        i++
                    // \x##
                    case 'x':
                        i+=3
                    default:
                        util.Panicf("Unknown escape: %v", c)
                }
        }
    }
    
    if str[0] == '"' {
        size--
    }
    if str[len(str)-1] == '"' {
        size--
    }
    
    return size
}
