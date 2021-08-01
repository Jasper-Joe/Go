func check(x interface{}) {
    res, ok := x.(float64)
    if ok {
        fmt.Println("This is a float", res)
    } else {
        fmt.Println("This is not a float", res)
    }
}


func switchCheck(x interface{}) {
    
    switch x.(type) {
        case int:
        fmt.Println("It is int")
        case int32:
        fmt.Println("It is int32")
        case int64:
        fmt.Println("It is int64")
        case float64:
        fmt.Println("It is float")
        case string:
        fmt.Println("It is string")
        default:
        fmt.Println("Custome type")
    }
}

func main() {
    var x interface{} = 88.8
    check(x)
    var a interface{} = 44
    var b interface{} = "Universally accessible"
    var c int64 = 90
    switchCheck(a)
    switchCheck(b)
    switchCheck(c)
}