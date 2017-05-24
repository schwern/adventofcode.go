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

func AssertStringSliceEq( t *testing.T, have, want []string ) {
    if len(have) != len(want) {
        t.Errorf( "\nlen(have): %v\nlen(want): %v", len(have), len(want) )
        return
    }
    
    for i := range have {
        if have[i] != want[i] {
            t.Errorf(
                "\nhave[%v]: %#v\nwant[%v]: %#v",
                i, have[i], i, want[i],
            )
            return
        }
    }
}
