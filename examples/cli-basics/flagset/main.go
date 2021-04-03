package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		os.Exit(2)
	}

	cmd := os.Args[1]

	switch cmd {
	case "greet":
		greetFlagSet := flag.NewFlagSet("greet", flag.ExitOnError)
		msgFlag := greetFlagSet.String("msg", "Welcome to Reminder CLI", "Welcome message to output when the executable is run")
		userFlag := greetFlagSet.String("user", "Tolu", "Name of user")
		err := greetFlagSet.Parse(os.Args[2:])

		if err != nil {
			log.Fatal("Flag Parse Error", err)
		}

		fmt.Printf("Hi %s\n", *userFlag)
		fmt.Printf("%s\n", *msgFlag)

	case "help":
		fmt.Println("Select help ID")
	default:
		fmt.Println("Invalid command provided")
	}
}
