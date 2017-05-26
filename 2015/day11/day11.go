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
    var pass string
    for pass = IncString(orig); !IsValidPassword(pass); pass = IncString(pass) {
    }
    
    return pass
}

func IncString( str string ) string {
    slice := []byte(str)
    
    for i := len(slice)-1; i >= 0; i-- {
        if slice[i] < 'z' {
            slice[i] = slice[i]+1
            break
        } else {
            slice[i] = 'a'
        }
    }
    
    return string(slice)
}
