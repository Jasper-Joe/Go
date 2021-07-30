func handlePanic() {
    if a := recover(); a != nil {
        fmt.Println("Recover", a)
    }
}

func helper(language *string, author *string) {
    defer fmt.Println("Defer from func")
    defer handlePanic()
    if language == nil {
        panic("Language cannot be empty")
    }
    
    if author == nil {
        panic("Author cannot be empty")
    }
    
    fmt.Println(*language, *author)
}

func main() {
    defer fmt.Println("Defer from main")
    l := "Eng"
    helper(&l, nil)
    fmt.Println("Out of main")
    
    /*
    Recover Author cannot be empty
    Defer from func
    Out of main
    Defer from main
    */
}