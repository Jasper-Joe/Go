type shape interface {
    area() int
    perimeter() int
}

type square struct {
    width int
    height int
}

func (s square) area() int {
    return s.width * s.height
}

func (s square) perimeter() int {
    return 2 * s.width + 2 * s.height
}


func main() {
    var sq shape
    sq = square{5, 10}
    res := sq.area()
    fmt.Println(res)
}