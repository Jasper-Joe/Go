package main

import (
	"flag"
	"fmt"
)

func main() {

	/*
		command line: go run say.go -orders=999 -fine=true
	*/
	fmt.Println("Starting application")
	fine := flag.Bool("fine", false, "fine flag")
	orders := flag.Int("orders", 1, "orders flag")
	flag.Parse()
	fmt.Println(*fine)
	fmt.Println(*orders)

}
