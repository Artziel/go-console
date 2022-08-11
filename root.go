package GoConsole

import (
	"fmt"
	"os"
	"regexp"
)

type Root struct {
	Init     func() error
	Commands map[string]Command
}

func (r *Root) Run() error {
	args := os.Args[1:]
	return r.run(args)
}

func (r *Root) run(args []string) error {
	var err error
	cmd := ""

	if len(args) > 0 {
		if args[0] == "help" || args[0] == "-h" {
			r.PrintHelp()
			return err
		}

		if match, _ := regexp.MatchString("^-(\\w)", args[0]); !match {
			cmd = args[0]
			args = args[1:]
		}
	}

	if _, found := r.Commands[cmd]; !found {
		return ErrSubCommandNotFound
	}

	command := r.Commands[cmd]
	if err = command.Execute(args); err != nil {
		return err
	}

	return nil
}

func (r *Root) help() string {
	help := ""
	for key, cmd := range r.Commands {
		if key == "" {
			help = help + fmt.Sprintf("- %s: ", Colored("<yellow>[With no SubCommand]</yellow>"))
		} else {
			help = help + fmt.Sprintf("- %s: ", Colored("<yellow>"+key+"</yellow>"))
		}
		help = help + fmt.Sprintf("%v\n", cmd.GetHelp())
	}

	return help
}

func (r *Root) PrintHelp() {
	fmt.Printf("%v", r.help())
}
