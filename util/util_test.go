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

func TestMustAtoi( t *testing.T ) {
    testutil.AssertEq( t, util.MustAtoi("-1234"), -1234 )
    testutil.AssertEq( t, util.MustAtoi("0"), 0 )

    defer testutil.AssertPanicf(t,
        `strconv.Atoi: parsing "one": invalid syntax`,
    )
    util.MustAtoi("one")
}

func TestMaxMinInts( t *testing.T ) {
    nums := []int{ 0, 10, 9999, -234 }
    
    testutil.AssertEq( t, util.MaxIntsIdx( nums ), 2 )
    testutil.AssertEq( t, util.MaxInts( nums ), 9999 )
    testutil.AssertEq( t, util.MinIntsIdx( nums ), 3 )
    testutil.AssertEq( t, util.MinInts( nums ), -234 )
}
