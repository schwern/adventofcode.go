package day04

import(
    "bytes"
    "crypto/md5"
    "encoding/hex"
    "strconv"
)

func MineAdventCoin( secret, prefix string ) int {    
    sumHex := make( []byte, md5.Size * 2 )
    for i := 0; i < 1e8; i++ {
        key := []byte( secret )
        key = strconv.AppendInt( key, int64(i), 10 )
        
        sum := md5.Sum(key)
        hex.Encode(sumHex, sum[:])

        if bytes.HasPrefix(sumHex, []byte(prefix)) {
            return i
        }
    }
    
    return -1
}
