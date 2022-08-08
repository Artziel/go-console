package GoConsole

import (
	"fmt"
	"strings"
)

func Hr(width int) {
	fmt.Printf("%v\n", strings.Repeat("-", width))
}
