package main

import (
	"fmt"
	"os"
	"time"

	GoConsole "github.com/artziel/go-console"
)

type SubCommandAFlags struct {
	String  string `GoConsole:"name:str,usage:String value example"`
	Boolean bool   `GoConsole:"name:b,usage:Boolean value example"`
	Integer int    `GoConsole:"name:i,usage:Integer value example"`
}

func main() {
	root := GoConsole.Root{
		Commands: map[string]GoConsole.Command{
			"": {
				Help:    "Main Command",
				Example: "$ sample",
				Flags: &SubCommandAFlags{
					String:  "test.yaml",
					Boolean: false,
					Integer: 123,
				},
				Run: func(args interface{}) error {
					flags := args.(*SubCommandAFlags)
					fmt.Printf("Main Command - Args: %v\n", flags)

					cmd := GoConsole.NewShellCommand("ls -la", "List files")
					if success := cmd.Exec(); !success {
						fmt.Printf("%v\n", cmd.GetLog())
					}

					spinner := GoConsole.Spinner09()
					spinner.SetColor("blue").SetPrefix("Test Prefix ").SetSufix(" Runing Process, please wait!").Start()

					time.Sleep(5 * time.Second)
					spinner.Stop()
					GoConsole.Printf("Process finish ... <green>OK!</green>\n")

					return nil
				},
			},
			"SubCommandA": {
				Help:    "SubCommand A",
				Example: "$ sample SubCommandA",
				Flags: &SubCommandAFlags{
					String:  "test.yaml",
					Boolean: false,
					Integer: 123,
				},
				Run: func(args interface{}) error {
					flags := args.(*SubCommandAFlags)
					fmt.Printf("SubCommand A - Args: %v\n", flags)
					return nil
				},
			},
		},
	}

	if err := root.Run(os.Args[1:]); err != nil {
		fmt.Println(err.Error())
	}
}
