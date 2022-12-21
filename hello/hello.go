package main

import (
    "fmt"
    "github.com/jy0ung/goproj/greet"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}