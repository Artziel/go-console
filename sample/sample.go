package main

import (
	"fmt"
	"os"

	GoConsole "github.com/artziel/go-console"
)

type SubCommandAFlags struct {
	String  string `GoConsole:"name:str,usage:String value example"`
	Boolean bool   `GoConsole:"name:b,usage:Boolean value example"`
	Integer int    `GoConsole:"name:i,usage:Integer value example"`
}

func main() {

	GoConsole.Printf("este es el <blue>%v</blue> y este es <red>%v</red> con variables\n", "AZUL", "ROJO")

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
		},
	}

	if err := root.Run(os.Args[1:]); err != nil {
		fmt.Println(err.Error())
	}
}
