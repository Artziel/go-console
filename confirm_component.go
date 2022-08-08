package GoConsole

import (
	"fmt"
	"strings"
)

type ConfirmOption struct {
	Label string
	Value string
}

type ConfirmComponent struct {
	Yes ConfirmOption
	No  ConfirmOption
}

func (c ConfirmComponent) Request() bool {
	options := fmt.Sprintf(Colored("<green>%v</green> / <yellow>%v</yellow>: "), c.Yes.Label, c.No.Label)
	fmt.Print(options)

	for {
		var input string
		fmt.Scanln(&input)
		value := strings.ToLower(input)

		if value == strings.ToLower(c.Yes.Value) {
			return true
		} else if value == strings.ToLower(c.No.Value) {
			return false
		} else {
			fmt.Print(options)
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
