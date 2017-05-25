package testutil

import(
    "fmt"
    "testing"
)

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

func AssertIntSliceEq( t *testing.T, have, want []int ) {
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

func AssertPanicf( t *testing.T, want_fmt string, want_args ...interface{} ) {
    want := fmt.Sprintf( want_fmt, want_args...)

    err := recover()
    
    if err == nil {
        t.Errorf("Expected panic: %v", want)
    } else {
        AssertEq( t, err, want )
    }
}
