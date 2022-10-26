package main

import "github.com/ZyperX/zbxss/internal/runner"
import "fmt"

func main() {
    options := runner.ParseOptions()
    fmt.Println(options)
    runner.New()
}
