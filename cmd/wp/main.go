package main

import (
	"fmt"
	"go-wp/internal/program"
)

var (
	version = "dev"
	date    = "unknown"
)

func main() {
	fmt.Printf("running wp %s, built at %s\n", version, date)
	program.Run()
}
