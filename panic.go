func helper(language *string, author *string) {
    defer fmt.Println("Defer from func")
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
    Defer from func
    Defer from main
    panic: Author cannot be empty
    */
}