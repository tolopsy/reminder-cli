package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		os.Exit(2)
	}

	cmd := os.Args[1]

	switch cmd {
	case "greet":
		msg := "Reminder CLI"
		if len(os.Args) > 2 {
			f := strings.Split(os.Args[2], "=")
			if len(f) == 2 && f[0] == "--msg" {
				msg = f[1]
			}
		}

		fmt.Printf("Welcome to %s", msg)

	case "help":
		fmt.Println("Select help ID")

	default:
		fmt.Println("Invalid command")
	}
}
