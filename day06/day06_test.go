package day06_test

import(
    "testing"
    "github.com/schwern/adventofcode2015/testutil"
    "github.com/schwern/adventofcode2015/util"
    "github.com/schwern/adventofcode2015/day06"
)

var Input_File = "../testdata/day06.txt"

func assertLightsEq( t *testing.T, have [][]bool, want [][]bool ) {
    if len(have) != len(want) {
        t.Errorf( "\nlen(have): %v\nlen(want): %v", have, want )
        return
    }
    
    for i := range want {
        if len(have[i]) != len(want[i]) {
            t.Errorf(
                "\nlen(have[%d]): %v\nlen(want[%d]): %v",
                i, have, i, want,
            )
            return
        }
        
        for j := range want[i] {
            if have[i][j] != want [i][j] {
                t.Errorf(
                    "\nhave[%d][%d]: %v\nwant[%d][%d]: %v",
                    i, j, have[i][j], i, j, want[i][j],
                )
                return
            }
        }
    }
}

func TestMakeLights( t *testing.T ) {
    lights := day06.MakeLights()
    testutil.AssertEq( t, len(lights), 1000 )
    testutil.AssertEq( t, len(lights[0]), 1000 )
    testutil.AssertEq( t, lights[0][0], false )
    testutil.AssertEq( t, lights[999][999], false )
}

func TestSwitchAllLightsOn( t *testing.T ) {
    have := day06.MakeLights()
    day06.SwitchLights( have, "turn on 0,0 through 999,999" )

    want := day06.MakeLights()
    for i := range want {
        for j := range want[i] {
            want[i][j] = true
        }
    }

    assertLightsEq( t, have, want )
}

func TestToggleLine( t *testing.T ) {
    have := day06.MakeLights()    
    day06.SwitchLights( have, "toggle 0,0 through 999,0" )
    
    want := day06.MakeLights()
    for i := 0; i < 1000; i++ {
        want[i][0] = true
    }
    
    assertLightsEq( t, have, want )
}

func TestPartOne( t *testing.T ) {
    lights := day06.MakeLights()
    
    lines := util.LineChannel( Input_File )
    for line := range lines {
        day06.SwitchLights( lights, line )
    }
    
    count := 0
    for i := range lights {
        for j := range lights[i] {
            if lights[i][j] {
                count++
            }
        }
    }
    
    testutil.AssertEq( t, count, 400410 )
}
