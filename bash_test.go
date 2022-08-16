package GoConsole

import "testing"

func TestNewShellCommand(t *testing.T) {

	c := NewShellCommand("ls -la ./samples", "Simple test List")

	parsed := c.parseCommand()

	if len(parsed) != 3 {
		t.Errorf("Test parseCommand unexpected result:\ngot  %v\nwant %v", parsed, []string{"ls", "-la", "./sample/*.*"})
	}

	c2 := NewShellCommand("ls -la ./samples/not-exists", "Simple test List")
	if ok := c2.Exec(); ok {
		t.Errorf("Test command execution fail: got\nsuccess\nwant: error")
	} else {
		l := c2.GetLog()
		expected := `Command "Simple test List" Execution FAIL
------------------------------------------------------------
ls: cannot access './samples/not-exists': No such file or directory

ERROR: exit status 2
`
		if l != expected {
			t.Errorf("Test command execution unexpected log: got  %s\nwant %s", l, expected)
		}
	}

}
