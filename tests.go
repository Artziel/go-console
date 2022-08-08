package GoConsole

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type testRootFlags struct {
	String  string `GoConsole:"name:str,usage:String value example"`
	Boolean bool   `GoConsole:"name:b,usage:Boolean value example"`
	Integer int    `GoConsole:"name:i,usage:Integer value example"`
}

var command = Command{
	Help:    "Main Command",
	Example: "$ sample",
	Flags: &testRootFlags{
		String:  "string test no subcommand",
		Boolean: false,
		Integer: 123,
	},
	Run: func(args interface{}) error {
		flags := args.(*testRootFlags)
		fmt.Printf("Main Command - Args: %v\n", flags)
		return nil
	},
}

var subCommand = Command{
	Help:    "SubCommand",
	Example: "$ sample SubCommand",
	Flags: &testRootFlags{
		String:  "string test subcommand",
		Boolean: false,
		Integer: 123,
	},
	Run: func(args interface{}) error {
		flags := args.(*testRootFlags)
		fmt.Printf("SubCommand- Args: %v\n", flags)
		return nil
	},
}

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}
