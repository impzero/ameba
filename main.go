package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/impzero/ameba/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello, %s! This is Ameba's official REPL", user.Username)
	fmt.Printf("\n")

	repl.Start(os.Stdin, os.Stdout)
}
