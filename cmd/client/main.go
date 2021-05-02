package main

import (
	"flag"
	"fmt"
	"go-projects/reminder-cli/client"
	"os"
)

var (
	backendURIFlag = flag.String("backend", "http://127.0.0.1:8080", "URI of backend API")
	helpFlag       = flag.Bool("help", false, "Display help message")
)

func main() {
	flag.Parse()
	s := client.NewSwitch(*backendURIFlag)

	if *helpFlag || len(os.Args) == 1 {
		s.Help()
		return
	}

	err := s.Switch()
	if err != nil {
		fmt.Printf("cmd switch error: %v\n", err)
		os.Exit(2)
	}
}
