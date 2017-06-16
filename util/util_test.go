package util_test

import(
    "runtime"
    "testing"
    "github.com/stvp/assert"
    "github.com/schwern/adventofcode.go/util"
)

func TestLineChannel( t *testing.T ) {
    _, file, _, _ := runtime.Caller(0)
    lines := util.LineChannel(file)
    
    assert.Equal( t, <-lines, "package util_test" )
    assert.Equal( t, <-lines, "" )
    for range lines {
        // do nothing, just making sure it ends
    }
    assert.Equal( t, true, true )
}

func TestMustAtoi( t *testing.T ) {
    assert.Equal( t, util.MustAtoi("-1234"), -1234 )
    assert.Equal( t, util.MustAtoi("0"), 0 )

    defer assert.Panic(t,
        `strconv.Atoi: parsing "one": invalid syntax`,
    )
    util.MustAtoi("one")
}

func TestMaxMinInts( t *testing.T ) {
    nums := []int{ 0, 10, 9999, -234 }
    
    assert.Equal( t, util.MaxIntsIdx( nums ), 2 )
    assert.Equal( t, util.MaxInts( nums ), 9999 )
    assert.Equal( t, util.MinIntsIdx( nums ), 3 )
    assert.Equal( t, util.MinInts( nums ), -234 )
}
