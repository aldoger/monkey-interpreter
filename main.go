package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/aldoger/monkey-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! this is the Monkey programming languange!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	fmt.Printf("Type exit to leave program\n")
	repl.Start(os.Stdin, os.Stdout)
}
