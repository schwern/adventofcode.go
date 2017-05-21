package day06

import(
    "regexp"
    "fmt"
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

func MakeLights() [][]bool {
    // Top level slice
    lights := make([][]bool, LightsY)
    
    // Slice for all pixels
    pixels := make([]bool, LightsX*LightsY)
    
    for i := range lights {
        lights[i], pixels = pixels[:LightsX], pixels[LightsX:]
    }
    
    return lights
}

func findAllNamed( re *regexp.Regexp, str string ) map[string]string {
    match := re.FindStringSubmatch( str )
    if match == nil {
        panic(fmt.Sprintf("Unknown instruction '%v'", str))
    }
    
    params := make(map[string]string)
    names := re.SubexpNames()
    for i := 1; i < len(names); i++ {
        params[names[i]] = match[i]
    }
    
    return params
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
    parsed := findAllNamed(instruction_re, instruction)
    
    cmd := OFF
    switch parsed["cmd"] {
        case "toggle":      cmd = TOGGLE
        case "turn on":     cmd = ON
        case "turn off":    cmd = OFF
        default:
            panic(fmt.Sprintf("Unknown command '%v'", parsed["cmd"]))
    }
    
    start := [2]int{ mustAtoi(parsed["x1"]), mustAtoi(parsed["y1"]) }
    end   := [2]int{ mustAtoi(parsed["x2"]), mustAtoi(parsed["y2"]) }
    
    return cmd, start, end
}

func SwitchLights( lights [][]bool, instruction string ) {
    cmd, start, end := parseInstruction( instruction )
    
    for x := start[0]; x <= end[0]; x++ {
        for y := start[1]; y <= end[1]; y++ {
            switch cmd {
                case ON:
                    lights[x][y] = true
                case OFF:
                    lights[x][y] = false
                case TOGGLE:
                    if lights[x][y] {
                        lights[x][y] = false
                    } else {
                        lights[x][y] = true
                    }
            }
        }
    }
}