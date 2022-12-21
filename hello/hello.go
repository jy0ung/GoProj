package main

import (
    "fmt"

    "https://github.com/jy0ung/goproj/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}