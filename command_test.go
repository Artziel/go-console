package GoConsole

import (
	"testing"
)

func TestCommand(t *testing.T) {

	output := captureOutput(func() {
		args := []string{
			"-str=test string",
			"-b",
			"-i=321",
		}
		err := command.Execute(args)
		if err != nil {
			t.Errorf("unexpected command error: %v", err.Error())
		}
	})

	expected := "Main Command - Args: &{test string true 321}\n"
	if output != expected {
		t.Errorf("unexpected command output: \n got %v\nwant %v", output, expected)
	}
}
