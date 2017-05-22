package util

import(
    "bufio"
    "os"
    "io/ioutil"
    "fmt"
)

func Check( err error ) {
    if err != nil {
        panic(err)
    }
}

func Panicf( format string, args ...interface{}) {
    panic(fmt.Sprintf( format, args... ))
}

func LineChannel( filename string ) (chan string) {
    ch := make( chan string )
    go func() {
        fh := OpenFile(filename)
        defer fh.Close()
        defer close(ch)

        scanner := bufio.NewScanner(fh)
        for scanner.Scan() {
            ch <- scanner.Text()
        }
    
        err := scanner.Err();
        Check(err)
    }()
    
    return ch
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
