package util

import(
    "regexp"
    "bufio"
    "os"
    "io/ioutil"
    "fmt"
)

func FindAllNamed( re *regexp.Regexp, str string ) map[string]string {
    match := re.FindStringSubmatch( str )
    if match == nil {
        panic("Match failed")
    }
    
    params := make(map[string]string)
    names := re.SubexpNames()
    for i := 1; i < len(names); i++ {
        params[names[i]] = match[i]
    }
    
    return params
}

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
