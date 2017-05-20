package day04

import(
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "strings"
)

func MineAdventCoin( secret string ) int {
    for i := 0; i < 1e8; i++ {
        key := []byte( fmt.Sprintf("%v%d", secret, i) )
        
        sum := md5.Sum(key)

        if strings.HasPrefix(hex.EncodeToString(sum[:]), "00000") {
            return i
        }
    }
    
    return -1
}
