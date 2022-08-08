package GoConsole

import (
	"fmt"
	"testing"
)

func TestRoot(t *testing.T) {

	root := Root{
		Commands: map[string]Command{
			"":           command,
			"subCommand": subCommand,
		},
	}

	output := captureOutput(func() {

		args := []string{
			"subCommand",
			"-str=test string",
			"-b",
			"-i=321",
		}
		if err := root.Run(args); err != nil {
			fmt.Println(err.Error())
		}
	})

	expected := "SubCommand- Args: &{test string true 321}\n"
	if output != expected {
		t.Errorf("unexpected command output: \n got %v\nwant %v", output, expected)
	}

	output2 := captureOutput(func() {

		args := []string{
			"-str=test string",
			"-b",
			"-i=321",
		}
		if err := root.Run(args); err != nil {
			fmt.Println(err.Error())
		}
	})

	expected = "Main Command - Args: &{test string true 321}\n"
	if output2 != expected {
		t.Errorf("unexpected command output: \n got %v\nwant %v", output2, expected)
	}
}
