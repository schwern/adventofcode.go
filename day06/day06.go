package day06

import(
    "regexp"
    "strconv"
    "github.com/schwern/adventofcode2015/util"
)

var LightsY = 1000
var LightsX = 1000

const(
    ON = iota
    OFF = iota
    TOGGLE = iota
)

func MakeLights() [][]int {
    // Top level slice
    lights := make([][]int, LightsY)
    
    // Slice for all pixels
    pixels := make([]int, LightsX*LightsY)
    
    for i := range lights {
        lights[i], pixels = pixels[:LightsX], pixels[LightsX:]
    }
    
    return lights
}

var instruction_re = regexp.MustCompile(
    `(?P<cmd>turn on|turn off|toggle) (?P<x1>\d+),(?P<y1>\d+) through (?P<x2>\d+),(?P<y2>\d+)`,
)

func mustAtoi( str string ) int {
    num, err := strconv.Atoi(str)
    util.Check(err)
    
    return num
}

func parseInstruction( instruction string ) (int, [2]int, [2]int) {
    parsed := util.FindAllNamed(instruction_re, instruction)
    
    cmd := OFF
    switch parsed["cmd"] {
        case "toggle":      cmd = TOGGLE
        case "turn on":     cmd = ON
        case "turn off":    cmd = OFF
        default:
            util.Panicf("Unknown command '%v'", parsed["cmd"])
    }
    
    start := [2]int{ mustAtoi(parsed["x1"]), mustAtoi(parsed["y1"]) }
    end   := [2]int{ mustAtoi(parsed["x2"]), mustAtoi(parsed["y2"]) }
    
    return cmd, start, end
}

func LightBright( lights [][]int, instruction string ) {
    cmd, start, end := parseInstruction( instruction )
    
    for x := start[0]; x <= end[0]; x++ {
        for y := start[1]; y <= end[1]; y++ {
            switch cmd {
                case ON:
                    lights[x][y]++
                case OFF:
                    if lights[x][y] > 0 {
                        lights[x][y]--
                    }
                case TOGGLE:
                    lights[x][y] += 2
            }
        }
    }
}

func SwitchLights( lights [][]int, instruction string ) {
    cmd, start, end := parseInstruction( instruction )
    
    for x := start[0]; x <= end[0]; x++ {
        for y := start[1]; y <= end[1]; y++ {
            switch cmd {
                case ON:
                    lights[x][y] = 1
                case OFF:
                    lights[x][y] = 0
                case TOGGLE:
                    if lights[x][y] != 0 {
                        lights[x][y] = 0
                    } else {
                        lights[x][y] = 1
                    }
            }
        }
    }
}