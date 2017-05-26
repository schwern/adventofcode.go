package day11

var validLength = 8
var pairsNeeded = 2

// Doing this char by char to do it in one pass of the string
// with early exit.
func IsValidPassword( pass string ) bool {
    // Exactly 8 characters
    if len(pass) != validLength {
        return false
    }
    
    nextPossiblePair := 0
    hasStraight := false
    numPairs := 0
    for i,char := range pass {
        // Not a lower case letter
        if char < 'a' || 'z' < char {
            return false
        }
        
        // Invalid chars
        switch char {
            case 'i', 'l', 'o':
                return false
        }
        
        // At least one increasing straight
        // of at least three letters,
        if !hasStraight {
            if i > validLength - 3 {
                return false
            } else {
                if pass[i+1] == pass[i]+1 && pass[i+2] == pass[i]+2 {
                    hasStraight = true
                }
            }
        }
        
        // At least two different, non-overlapping pairs 
        if numPairs < pairsNeeded {
            if i > validLength - 2 {
                return false
            } else {
                if i >= nextPossiblePair && pass[i] == pass[i+1] {
                    numPairs++
                    nextPossiblePair = i+2
                }
            }
        }
    }
    
    return hasStraight && numPairs >= pairsNeeded
}

func NextPassword( orig string ) string {
    pass := make( []byte, len(orig) )
    copy( pass, orig )
    
    for IncByteSlice(pass); !IsValidPassword(string(pass)); IncByteSlice(pass) {
    }
    
    return string(pass)
}

func IncByteSlice( str []byte ) {
    for i := len(str)-1; i >= 0; i-- {
        if str[i] < 'z' {
            str[i] = str[i]+1
            return
        } else {
            str[i] = 'a'
        }
    }
}
