package client

import (
	"fmt"
	"os"
)

type BackendHTTPClient interface {
}

type Switch struct {
	client        BackendHTTPClient
	backendAPIURL string
	commands      map[string]func() func(string) error
}

func NewSwitch(uri string) Switch {
	httpClient := NewHTTPClient(uri)
	commandSwitch := Switch{
		client:        httpClient,
		backendAPIURL: uri,
	}
	commandSwitch.commands = map[string]func() func(string) error{
		"create": commandSwitch.create,
		"edit":   commandSwitch.edit,
		"fetch":  commandSwitch.fetch,
		"delete": commandSwitch.delete,
		"health": commandSwitch.health,
	}
	return commandSwitch
}

func (s Switch) Switch() error {
	cmdName := os.Args[1]
	cmd, ok := s.commands[cmdName]

	if !ok {
		return fmt.Errorf("Invalid command '%s'", cmdName)
	}

	return cmd()(cmdName)
}

func (s Switch) create() func(string) error {
	return func(s string) error {
		fmt.Println("Create Reminder")
		return nil
	}
}

func (s Switch) edit() func(string) error {
	return func(s string) error {
		fmt.Println("Edit Reminder")
		return nil
	}
}

func (s Switch) fetch() func(string) error {
	return func(s string) error {
		fmt.Println("Fetch Reminder")
		return nil
	}
}

func (s Switch) delete() func(string) error {
	return func(s string) error {
		fmt.Println("Delete Reminder")
		return nil
	}
}

func (s Switch) health() func(string) error {
	return func(s string) error {
		fmt.Println("Reminder CLI Health")
		return nil
	}
}
