package mapWithDefault_test

import(
    "testing"
    "github.com/schwern/adventofcode.go/mapWithDefault"
    "github.com/schwern/adventofcode.go/testutil"
)

func TestMapWithDefault( t *testing.T ) {
    dflt := 9999
    m := mapWithDefault.New( dflt )
    testutil.AssertEq( t, m.Get("I do not exist"), dflt )
    testutil.AssertEq( t, m.Exists("I do not exist"), false )
    
    m.Set("this", "that")
    testutil.AssertEq( t, m.Get("this"), "that" )
    testutil.AssertEq( t, m.Exists("this"), true )
    
    m.Delete("this")
    testutil.AssertEq( t, m.Get("this"), dflt )
    testutil.AssertEq( t, m.Exists("this"), false )
}
