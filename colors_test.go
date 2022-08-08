package GoConsole

import (
	"testing"
)

type colorTests struct {
	color    string
	input    string
	expected string
}

func TestColors(t *testing.T) {
	tests := []colorTests{
		{color: "RED", input: Colored("<red>RED</red>"), expected: "\033[31mRED\033[0m"},
		{color: "GREEN", input: Colored("<green>GREEN</green>"), expected: "\033[32mGREEN\033[0m"},
		{color: "YELLOW", input: Colored("<yellow>YELLOW</yellow>"), expected: "\033[33mYELLOW\033[0m"},
		{color: "BLUE", input: Colored("<blue>BLUE</blue>"), expected: "\033[34mBLUE\033[0m"},
		{color: "PURPLE", input: Colored("<purple>PURPLE</purple>"), expected: "\033[35mPURPLE\033[0m"},
		{color: "CYAN", input: Colored("<cyan>CYAN</cyan>"), expected: "\033[36mCYAN\033[0m"},
		{color: "WHITE", input: Colored("<white>WHITE</white>"), expected: "\033[37mWHITE\033[0m"},
	}

	for i, test := range tests {
		if test.input != test.expected {
			t.Errorf("Test %d: Text will not print as expected (color %s) \n", i, test.color)
		}
	}
}
