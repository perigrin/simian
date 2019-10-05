package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/perigrin/simian/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hi %s! This is the Simian langauge!\n", user.Username)
	fmt.Printf("Feel free to type in commands.\n")
	fmt.Printf("(Use Ctrl-D to stop)\n")
	repl.Start(os.Stdin, os.Stdout)
}
