package GoConsole

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func ScanString() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	return text
}

func ScanInt() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	if v, err := strconv.Atoi(input); err == nil {
		return v
	}

	return 0
}

func WaitForEnter(msg string) {
	fmt.Println(msg)
	var input string
	fmt.Scanln(&input)
}

func windowsClearTerm() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func linuxClearTerm() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Clear() {
	cmds := map[string]func(){
		"windows": windowsClearTerm,
		"linux":   linuxClearTerm,
	}

	value, ok := cmds[runtime.GOOS]
	if ok {
		value()
	}
}
