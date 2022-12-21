package main

import (
    "fmt"
	"github.com/google/go-github/v48/github"
    "github.com/jy0ung/goproj/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}