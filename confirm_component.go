package GoConsole

import (
	"errors"
	"strings"
)

var ErrConfirmInvalidOption error = errors.New("invalid input option")

type ConfirmOption struct {
	Label string
	Value string
}

type ConfirmComponent struct {
	Yes ConfirmOption
	No  ConfirmOption
}

func (c *ConfirmComponent) printOptions() {
	Printf("<green>%s</green> / <yellow>%s</yellow>: ", c.Yes.Label, c.No.Label)
}

func (c *ConfirmComponent) eval(value string) (bool, error) {
	var err error
	value = strings.ToLower(value)
	result := false

	if value != strings.ToLower(c.Yes.Value) && value != strings.ToLower(c.No.Value) {
		err = ErrConfirmInvalidOption
	}

	if value == strings.ToLower(c.Yes.Value) {
		result = true
	}

	if value == strings.ToLower(c.No.Value) {
		result = false
	}

	return result, err
}

func (c *ConfirmComponent) Request() bool {

	c.printOptions()

	for {
		input := ScanString()
		val, err := c.eval(input)

		if err != nil {
			c.printOptions()
		} else {
			return val
		}
	}
}

func NewConfirm(yes ConfirmOption, no ConfirmOption) ConfirmComponent {
	c := ConfirmComponent{
		Yes: yes,
		No:  no,
	}

	return c
}
