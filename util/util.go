package util

import(
    "os"
    "io/ioutil"
)

func Check( err error ) {
    if err != nil {
        panic(err)
    }
}

func OpenFile( filename string ) *os.File {
    fh, err := os.Open(filename)
    Check(err)
    
    return fh
}

func ReadFile( filename string ) string {
    content, err := ioutil.ReadFile( filename )
    Check(err)
    
    return string(content)
}