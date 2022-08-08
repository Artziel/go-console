package GoConsole

import (
	"fmt"
	"strings"

	Rainbow "gitlab.com/alus/go-console-colors"
)

func Hr() {
	fmt.Printf("%v\n", strings.Repeat("-", 60))
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

	return Rainbow.Colored(result)
}

func PrintParagraph(str string, width int) {
	fmt.Print(FormatParagraph(str, 60))
}

func Println(str string) {
	fmt.Println(Rainbow.Colored(str))
}
