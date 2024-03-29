package main

import (
	"fmt"
	"os"
	"os/user"

	"langInterpreter/repl"
)

func main() {
	userAcc, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", userAcc.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
