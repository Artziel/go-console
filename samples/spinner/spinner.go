package main

import (
	"fmt"
	"time"

	GoConsole "github.com/artziel/go-console"
)

func main() {
	spinners := []*GoConsole.Spinner{
		GoConsole.Spinner01(),
		GoConsole.Spinner02(),
		GoConsole.Spinner03(),
		GoConsole.Spinner04(),
		GoConsole.Spinner05(),
		GoConsole.Spinner06(),
		GoConsole.Spinner07(),
		GoConsole.Spinner08(),
		GoConsole.Spinner09(),
	}

	for i, spinner := range spinners {
		spinner.SetColor("blue").SetPrefix(fmt.Sprintf("Spinner %d - Prefix ", i+1)).SetSufix(" Spinner Sufix").Start()
		time.Sleep(3 * time.Second)
		spinner.Stop()
		fmt.Printf("Spinner %d Stoped!\n", i+1)
	}
}
