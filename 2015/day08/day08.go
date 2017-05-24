package day08

import "unicode"

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
