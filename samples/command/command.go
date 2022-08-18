package main

import (
	"fmt"

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
			"SubCommandB": {
				Help:    "SubCommand B",
				Example: "$ sample SubCommandB",
				Run: func(args interface{}) error {
					fmt.Println("SubCommand B - without flags")
					return nil
				},
			},
		},
	}

	if err := root.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
