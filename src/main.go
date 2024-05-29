package main

import (
	"fmt"
	"os"
	"os/user"
	"reboot/repl"
	"reboot/runner"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is Reboot programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)

		return
	}
	pathToFile := os.Args[1]

	if strings.Split(pathToFile, ".")[1] != "re" {
		panic("invalid file")
	}
	runner.Run(pathToFile)

}
