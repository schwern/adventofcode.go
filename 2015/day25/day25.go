package day25

func NextCode( code int ) int {
    return (code * 252533) % 33554393
}

func sumTo( n int ) (sum int) {
    for i := 1; i <= n; i++ {
        sum += i
    }
    return
}

func GridPos( row, col int ) int {
    pos := sumTo(col)
    
    for i := 2; i <= row; i++ {
        pos += col + i - 2
    }
    
    return pos
}
