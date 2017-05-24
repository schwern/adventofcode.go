package util_test

import(
    "runtime"
    "testing"
    "github.com/schwern/adventofcode.go/testutil"
    "github.com/schwern/adventofcode.go/util"
)

func TestLineChannel( t *testing.T ) {
    _, file, _, _ := runtime.Caller(0)
    lines := util.LineChannel(file)
    
    testutil.AssertEq( t, <-lines, "package util_test" )
    testutil.AssertEq( t, <-lines, "" )
    for range lines {
        // do nothing, just making sure it ends
    }
    testutil.AssertEq( t, true, true )
}
