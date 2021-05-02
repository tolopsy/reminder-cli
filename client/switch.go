package client

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type idsFlag []string

func (list idsFlag) String() string {
	return strings.Join(list, ",")
}

func (list *idsFlag) Set(s string) error {
	*list = append(*list, s)
	return nil
}

type BackendHTTPClient interface {
	Create(title, message string, duration time.Duration) ([]byte, error)
	Edit(id, title, message string, duration time.Duration) ([]byte, error)
	Fetch(ids []string) ([]byte, error)
	Delete(ids []string) error
	HealthCheck(host string) bool
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

func (s Switch) Help() {
	var help string
	for name := range s.commands {
		help += name + "\t --help\n"
	}
	fmt.Printf("Usage of %s:\n <command> [<args>]\n%s", os.Args[0], help)
}

func (s Switch) create() func(string) error {
	return func(cmd string) error {
		createCommand := flag.NewFlagSet(cmd, flag.ExitOnError)
		t, m, d := s.reminderFlags(createCommand)

		if err := s.checkArgs(3); err != nil {
			return err
		}

		if err := s.parseCmd(createCommand); err != nil {
			return err
		}

		res, err := s.client.Create(*t, *m, *d)
		if err != nil {
			return wrapError("Could not create reminder", err)
		}

		fmt.Printf("Reminder created successfully:\n%s", string(res))
		return nil
	}
}

func (s Switch) edit() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		editCommand := flag.NewFlagSet(cmd, flag.ExitOnError)
		editCommand.Var(&ids, "id", "The ID (int) of the reminder to edit")
		t, m, d := s.reminderFlags(editCommand)

		if err := s.checkArgs(2); err != nil {
			return err
		}

		if err := s.parseCmd(editCommand); err != nil {
			return err
		}

		lastID := ids[len(ids)-1]
		res, err := s.client.Edit(lastID, *t, *m, *d)
		if err != nil {
			return wrapError("Could not edit reminder", err)
		}

		fmt.Printf("Reminder edited successfully:\n%s", string(res))
		return nil
	}
}

func (s Switch) fetch() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		fetchCommand := flag.NewFlagSet(cmd, flag.ExitOnError)
		fetchCommand.Var(&ids, "id", "List of reminder IDs (int) to fetch")

		if err := s.checkArgs(1); err != nil {
			return err
		}

		if err := s.parseCmd(fetchCommand); err != nil {
			return err
		}

		res, err := s.client.Fetch(ids)
		if err != nil {
			return wrapError("Could not fetch reminder(s)", err)
		}

		fmt.Printf("Reminders fetched successfully:\n%s", string(res))
		return nil
	}
}

func (s Switch) delete() func(string) error {
	return func(cmd string) error {
		ids := idsFlag{}
		deleteCommand := flag.NewFlagSet(cmd, flag.ExitOnError)
		deleteCommand.Var(&ids, "id", "List of reminder IDs (int) to delete")

		if err := s.checkArgs(1); err != nil {
			return err
		}

		if err := s.parseCmd(deleteCommand); err != nil {
			return err
		}

		err := s.client.Delete(ids)
		if err != nil {
			return wrapError("Could not delete reminder(s)", err)
		}

		fmt.Printf("Reminder(s) deleted successfully:\n%v\n", ids)
		return nil
	}
}

func (s Switch) health() func(string) error {
	return func(cmd string) error {
		var host string
		healthCommand := flag.NewFlagSet(cmd, flag.ExitOnError)
		healthCommand.StringVar(&host, "host", s.backendAPIURL, "Host to call for health")

		if err := s.parseCmd(healthCommand); err != nil {
			return err
		}

		if !s.client.HealthCheck(host) {
			fmt.Printf("Host %s is down\n", host)
		} else {
			fmt.Printf("Host %s is up and running\n", host)
		}

		return nil
	}
}

func (s Switch) reminderFlags(f *flag.FlagSet) (*string, *string, *time.Duration) {
	t, m, d := "", "", time.Duration(0)
	f.StringVar(&t, "title", "", "Title of Reminder")
	f.StringVar(&t, "t", "", "Title of Reminder")
	f.StringVar(&m, "Message", "", "Message of Reminder")
	f.StringVar(&m, "m", "", "Message of Reminder")
	f.DurationVar(&d, "Duration", 0, "Duration of Reminder")
	f.DurationVar(&d, "d", 0, "Duration of Reminder")
	return &t, &m, &d
}

func (s Switch) checkArgs(minimumArgs int) error {
	if len(os.Args) == 3 && os.Args[2] == "--help" {
		return nil
	}

	if len(os.Args)-2 < minimumArgs {
		fmt.Printf("Incorrect use of %s\n %s %s --help\n", os.Args[1], os.Args[0], os.Args[1])
		return fmt.Errorf("%s expects at least %d arg(s), %d  provided", os.Args[1], minimumArgs, len(os.Args)-2)
	}

	return nil
}

func (s Switch) parseCmd(cmd *flag.FlagSet) error {
	err := cmd.Parse(os.Args[2:])
	if err != nil {
		return wrapError("Could not parse '"+cmd.Name()+"' command flags", err)
	}
	return nil
}
