# Golang Console 
Artziel Narvaiza <artziel@gmail.com>

### Features
- Allow use of sub command

### Dependencies
- gitlab.com/alus/go-console-colors

Use example:
```golang
package main

import (
	"fmt"
	"os"

	GoConsole "gitlab.com/Artziel/go-console"
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
		},
	}

	if err := root.Run(os.Args[1:]); err != nil {
		fmt.Println(err.Error())
	}
}
```
