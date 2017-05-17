package testutil_test

import(
    "github.com/schwern/testutil"
    "testing"
)

func TestAssertEq( t *testing.T ) {
    testutil.AssertEq( t, 10, 10 )
    testutil.AssertEq( t, "foo", "foo" )
}
