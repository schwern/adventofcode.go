package testutil_test

import(
    "github.com/schwern/adventofcode.go/testutil"
    "testing"
)

func TestAssertEq( t *testing.T ) {
    testutil.AssertEq( t, 10, 10 )
    testutil.AssertEq( t, "foo", "foo" )
}

func TestAssertStringSliceEq( t *testing.T ) {
    testutil.AssertStringSliceEq( t, []string{}, []string{} )
    testutil.AssertStringSliceEq( t,
        []string{"foo", "bar"},
        []string{"foo", "bar"},
    )
}

func TestAssertIntSliceEq( t *testing.T ) {
    testutil.AssertIntSliceEq( t, []int{}, []int{} )
    testutil.AssertIntSliceEq( t,
        []int{2, 5},
        []int{2, 5},
    )
}

func TestAssertPanicf( t *testing.T ) {
    defer testutil.AssertPanicf( t, "This: %v", 42 )
    panic("This: 42")
}
