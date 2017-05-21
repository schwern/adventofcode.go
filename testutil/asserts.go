package testutil

import "testing"

type TestCase struct {
    Arg string
    Want int
}

func AssertEq(t *testing.T, have, want interface{}) {
    if have != want {
        t.Errorf(
            "\nhave: %#v\nwant: %#v",
            have, want,
        )
    }
}
