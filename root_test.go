package GoConsole

import (
	"fmt"
	"testing"
)

const testCommandHelp = `
- <yellow>[With no SubCommand]</yellow>: Main Command

    Example: $ sample

    <blue>-str</blue>: String value example
    <blue>-b</blue>: Boolean value example
    <blue>-i</blue>: Integer value example

- <yellow>SubCommandA</yellow>: SubCommand A

    Example: $ sample SubCommandA

    <blue>-str</blue>: String value example
    <blue>-b</blue>: Boolean value example
    <blue>-i</blue>: Integer value example

`

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
		if err := root.run(args); err != nil {
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
		if err := root.run(args); err != nil {
			fmt.Println(err.Error())
		}
	})

	expected = "Main Command - Args: &{test string true 321}\n"
	if output2 != expected {
		t.Errorf("unexpected command output: \n got %v\nwant %v", output2, expected)
	}

	h := root.help()
	expected = Colored(testCommandHelp)
	if h != expected {
		t.Errorf("unexpected command help output: \ngot\n%s\nwant\n%s", h, expected)
	}

}
