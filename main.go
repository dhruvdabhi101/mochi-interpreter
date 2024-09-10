package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/dhruvdabhi101/interpreter/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Mochi Programming language! Learn with Dhruv!\n", user.Username)

	fmt.Println("Feel free to type in Commands")
	repl.Start(os.Stdin, os.Stdout)
}
