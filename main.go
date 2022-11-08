package main

import (
	"fmt"
	"github.com/borghippo/monkey-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to Monkey!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
