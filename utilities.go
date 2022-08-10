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
	Println(msg)
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

func formatToWidth(str string, width int) string {
	result := ""

	words := strings.Split(str, " ")
	line := ""
	for _, word := range words {
		if len([]rune(line))+len([]rune(word))+1 > width {
			result = result + "\n" + line
			line = word + " "
		} else {
			line = line + word + " "
		}
	}

	result = result + "\n" + line

	return result
}

func FormatParagraph(str string, width int) string {
	result := ""

	ps := strings.Split(str, "\n")

	for _, p := range ps {
		result = result + formatToWidth(p, width)
	}

	return result
}

func PrintParagraph(str string, width int) {
	Print(FormatParagraph(str, width))
}

func Print(str string) {
	fmt.Print(Colored(str))
}

func Println(str string) {
	fmt.Println(Colored(str))
}

func Printf(format string, values ...interface{}) {
	fmt.Printf(Colored(format), values...)
}
