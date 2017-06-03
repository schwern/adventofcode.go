package day18

type Grid [][]bool

type GOL struct {
    grid Grid
    x int
    y int
}

func (self *GOL) Grid() Grid {
    return self.grid
}

func (self *GOL) makeGrid() Grid {
    // Top level slice
    grid := make(Grid, self.y)
        
    // Slice for all lights
    lights := make([]bool, self.x * self.y)
        
    for i := range grid {
        grid[i], lights = lights[:self.x], lights[self.x:]
    }
        
    return grid
}

func (self *GOL) nextLightState( lightX int, lightY int ) bool {
    neighbors := 0    
    for x := lightX - 1; x <= (lightX + 1); x++ {
        if x < 0 || self.x <= x {
            continue
        }
        for y := lightY - 1; y <= (lightY + 1); y++ {
            if y < 0 || self.y <= y {
                continue
            }
            if x == lightX && y == lightY {
                continue
            }
            
            if self.grid[x][y] {
                neighbors++
            }
        }
    }

    switch self.grid[lightX][lightY] {
        case true:
            if 2 <= neighbors && neighbors <= 3 {
                return true
            } else {
                return false
            }
        case false:
            if neighbors == 3 {
                return true
            } else {
                return false
            }
    }
    
    return false
}

func (self *GOL) nextGrid() Grid {
    nextGrid := self.makeGrid()
    
    for x := range self.grid {
        for y := range self.grid[x] {
            nextGrid[x][y] = self.nextLightState(x,y)
        }
    }
    
    return nextGrid
}

func (self *GOL) Step() {
    self.grid = self.nextGrid()
}

func (self *GOL) HowManyLightsDoYouSee() int {
    numLights := 0
    
    for x := range self.grid {
        for y := range self.grid[x] {
            if self.grid[x][y] {
                numLights++
            }
        }
    }
    
    return numLights
}

func NewGOL( x int, y int, state Grid ) *GOL {    
    gol := GOL{ x: x, y: y, grid: state }

    return &gol
}
