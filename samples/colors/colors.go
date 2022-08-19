package main

import (
	"fmt"

	GoConsole "github.com/artziel/go-console"
)

func main() {
	output := []string{
		GoConsole.Colored("<red>RED   </red>\t<red bold>RED Bold</red>\t<red underline>RED Underline</red>    \t<red bold underline>RED Bold Underline</red>"),
		GoConsole.Colored("<green>GREEN</green>\t<green bold>GREEN Bold</green>\t<green underline>GREEN Underline</green>  \t<green bold underline>GREEN Bold Underline</green>"),
		GoConsole.Colored("<yellow>YELLOW</yellow>\t<yellow bold>YELLOW Bold</yellow>\t<yellow underline>YELLOW Underline</yellow>\t<yellow bold underline>YELLOE Bold Underline</yellow>"),
		GoConsole.Colored("<blue>BLUE </blue>\t<blue bold>BLUE Bold</blue>\t<blue underline>BLUE Underline</blue>  \t<blue bold underline>BLUE Bold Underline</blue>"),
		GoConsole.Colored("<purple>PURPLE</purple>\t<purple bold>PURPLE Bold</purple>\t<purple underline>PURPLE Underline</purple>\t<purple bold underline>PURPLE Bold Underline</purple>"),
		GoConsole.Colored("<cyan>CYAN </cyan>\t<cyan bold>CYAN Bold</cyan>\t<cyan underline>CYAN Underline</cyan>  \t<cyan bold underline>CYAN Bold Underline</cyan>"),
		GoConsole.Colored("<white>WHITE</white>\t<white bold>WHITE Bold</white>\t<white underline>WHITE Underline</white>  \t<white bold underline>WHITE Bold Underline</white>"),
	}

	for i, ln := range output {
		fmt.Printf("%d) %s\n", i+1, ln)
	}
}
