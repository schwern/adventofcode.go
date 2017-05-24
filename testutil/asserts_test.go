package testutil_test

import(
    "github.com/schwern/adventofcode.go/testutil"
    "testing"
)

func TestAssertEq( t *testing.T ) {
    testutil.AssertEq( t, 10, 10 )
    testutil.AssertEq( t, "foo", "foo" )
}

func TestAssertPanicf( t *testing.T ) {
    defer testutil.AssertPanicf( t, "This: %v", 42 )
    panic("This: 42")
}
