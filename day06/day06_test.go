package day06_test

import(
    "testing"
    "github.com/schwern/adventofcode2015/testutil"
    "github.com/schwern/adventofcode2015/util"
    "github.com/schwern/adventofcode2015/day06"
)

var Input_File = "../testdata/day06.txt"

func assertLightsEq( t *testing.T, have [][]int, want [][]int ) {
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
    testutil.AssertEq( t, lights[0][0], 0 )
    testutil.AssertEq( t, lights[999][999], 0 )
}

func TestSwitchAllLightsOn( t *testing.T ) {
    have := day06.MakeLights()
    day06.SwitchLights( have, "turn on 0,0 through 999,999" )

    want := day06.MakeLights()
    for i := range want {
        for j := range want[i] {
            want[i][j] = 1
        }
    }

    assertLightsEq( t, have, want )
}

func TestToggleLine( t *testing.T ) {
    have := day06.MakeLights()    
    day06.SwitchLights( have, "toggle 0,0 through 999,0" )
    
    want := day06.MakeLights()
    for i := 0; i < 1000; i++ {
        want[i][0] = 1
    }
    
    assertLightsEq( t, have, want )
}

func totalOn( lights [][]int ) int {
    count := 0
    for i := range lights {
        for j := range lights[i] {
            if lights[i][j] != 0 {
                count++
            }
        }
    }
    
    return count
}

func totalBrightness( lights [][]int ) int {
    brightness := 0
    for i := range lights {
        for j := range lights[i] {
            brightness += lights[i][j]
        }
    }
    
    return brightness
}

func TestPartOne( t *testing.T ) {
    lights := day06.MakeLights()
    
    lines := util.LineChannel( Input_File )
    for line := range lines {
        day06.SwitchLights( lights, line )
    }
    
    testutil.AssertEq( t, totalOn(lights), 400410 )
}

func TestLightBright( t *testing.T ) {
    have := day06.MakeLights()
    day06.LightBright( have, "turn on 0,0 through 0,0" )
    
    want := day06.MakeLights()
    want[0][0] = 1
    assertLightsEq( t, have, want )
}

func TestPartTwo( t *testing.T ) {
    lights := day06.MakeLights()
    
    lines := util.LineChannel( Input_File )
    for line := range lines {
        day06.LightBright( lights, line )
    }
    
    testutil.AssertEq( t, totalBrightness( lights ), 15343601 )
}
