package day04

import(
    "crypto/md5"
    "encoding/hex"
    "strconv"
    "strings"
)

func MineAdventCoin( secret, prefix string ) int {
    for i := 0; i < 1e8; i++ {
        key := []byte( secret )
        key = strconv.AppendInt( key, int64(i), 10 )
        
        sum := md5.Sum(key)

        if strings.HasPrefix(hex.EncodeToString(sum[:]), prefix) {
            return i
        }
    }
    
    return -1
}
