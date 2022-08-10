package GoConsole

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type ShellCommand struct {
	// Public
	Description string

	// Private
	log string
	cmd string
}

func (c *ShellCommand) GetLog() string {
	return c.log
}

func (c *ShellCommand) parseCommand() []string {
	quoted := false
	a := strings.FieldsFunc(c.cmd, func(r rune) bool {
		if r == '"' {
			quoted = !quoted
		}
		return !quoted && r == ' '
	})
	return a
}

func (c *ShellCommand) Exec() bool {
	var (
		outb, errb bytes.Buffer
		success    bool = true
		cmd        *exec.Cmd
	)

	execCMD := c.parseCommand()

	if len(execCMD) > 1 {
		cmd = exec.Command(execCMD[0], execCMD[1:]...)
	} else {
		cmd = exec.Command(execCMD[0])
	}

	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		success = false
		c.log = c.log + fmt.Sprintf("Command \"%s\" Execution FAIL\n%s\n%s\n", c.Description, strings.Repeat("-", 60), errb.String())
	}

	return success
}

func NewShellCommand(cmd string, description string) ShellCommand {
	return ShellCommand{
		Description: cmd,
		cmd:         description,
	}
}
