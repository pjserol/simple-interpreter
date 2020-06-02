package main

import (
	"fmt"
	"os"
	"os/user"
	"pjserol/simple-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is my first interpreter programming language!\n", user.Username)
	fmt.Printf("Feel free to type what you want...\n")
	repl.Start(os.Stdin, os.Stdout)
}
