import(
    "fmt"
    "time"
)

func iter(s string) {
    for i := 0; i < 6; i++ {
        time.Sleep(1 * time.Second)
        fmt.Println(s, i)
    }
}

func main() {
    go iter("New Goroutine")
    iter("Normoal Call")
    fmt.Printf("Done")
    
}

/*
New Goroutine 0
Normoal Call 0
Normoal Call 1
New Goroutine 1
New Goroutine 2
Normoal Call 2
Normoal Call 3
New Goroutine 3
New Goroutine 4
Normoal Call 4
New Goroutine 5
Normoal Call 5
Done
*/