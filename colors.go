package GoConsole

import (
	"regexp"
	"strings"
)

const (
	Reset = "\033[0m"

	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"

	RedUnderlined    = "\033[4;31m"
	GreenUnderlined  = "\033[4;32m"
	YellowUnderlined = "\033[4;33m"
	BlueUnderlined   = "\033[4;34m"
	PurpleUnderlined = "\033[4;35m"
	CyanUnderlined   = "\033[4;36m"
	WhiteUnderlined  = "\033[4;37m"

	RedBoldUnderlined    = "\033[1;4;31m"
	GreenBoldUnderlined  = "\033[1;4;32m"
	YellowBoldUnderlined = "\033[1;4;33m"
	BlueBoldUnderlined   = "\033[1;4;34m"
	PurpleBoldUnderlined = "\033[1;4;35m"
	CyanBoldUnderlined   = "\033[1;4;36m"
	WhiteBoldUnderlined  = "\033[1;4;37m"
)

func Colored(str string) string {
	tags := []string{
		"error",
		"warning",
		"info",
		"success",

		"green", "green underline", "green bold", "green bold underline", "green underline bold",
		"red", "red underline", "red bold", "red bold underline", "red underline bold",
		"cyan", "cyan underline", "cyan bold", "cyan bold underline", "cyan underline bold",
		"yellow", "yellow underline", "yellow bold", "yellow bold underline", "yellow underline bold",
		"blue", "blue underline", "blue bold", "blue bold underline", "blue underline bold",
		"purple", "purple underline", "purple bold", "purple bold underline", "purple underline bold",
		"white", "white underline", "white bold", "white bold underline", "white underline bold",
	}
	formatted := str

	for _, t := range tags {
		switch t {
		case "error":
			formatted = strings.Replace(formatted, "<"+t+">", Red, -1)
		case "warning":
			formatted = strings.Replace(formatted, "<"+t+">", Yellow, -1)
		case "info":
			formatted = strings.Replace(formatted, "<"+t+">", Blue, -1)
		case "success":
			formatted = strings.Replace(formatted, "<"+t+">", Green, -1)

		case "red":
			formatted = strings.Replace(formatted, "<"+t+">", Red, -1)
		case "green":
			formatted = strings.Replace(formatted, "<"+t+">", Green, -1)
		case "yellow":
			formatted = strings.Replace(formatted, "<"+t+">", Yellow, -1)
		case "cyan":
			formatted = strings.Replace(formatted, "<"+t+">", Cyan, -1)
		case "blue":
			formatted = strings.Replace(formatted, "<"+t+">", Blue, -1)
		case "purple":
			formatted = strings.Replace(formatted, "<"+t+">", Purple, -1)
		case "white":
			formatted = strings.Replace(formatted, "<"+t+">", White, -1)

		case "red underline":
			formatted = strings.Replace(formatted, "<"+t+">", RedUnderlined, -1)
		case "green underline":
			formatted = strings.Replace(formatted, "<"+t+">", GreenUnderlined, -1)
		case "yellow underline":
			formatted = strings.Replace(formatted, "<"+t+">", YellowUnderlined, -1)
		case "cyan underline":
			formatted = strings.Replace(formatted, "<"+t+">", CyanUnderlined, -1)
		case "blue underline":
			formatted = strings.Replace(formatted, "<"+t+">", BlueUnderlined, -1)
		case "purple underline":
			formatted = strings.Replace(formatted, "<"+t+">", PurpleUnderlined, -1)
		case "white underline":
			formatted = strings.Replace(formatted, "<"+t+">", WhiteUnderlined, -1)

		case "red bold":
			formatted = strings.Replace(formatted, "<"+t+">", RedUnderlined, -1)
		case "green bold":
			formatted = strings.Replace(formatted, "<"+t+">", GreenUnderlined, -1)
		case "yellow bold":
			formatted = strings.Replace(formatted, "<"+t+">", YellowUnderlined, -1)
		case "cyan bold":
			formatted = strings.Replace(formatted, "<"+t+">", CyanUnderlined, -1)
		case "blue bold":
			formatted = strings.Replace(formatted, "<"+t+">", BlueUnderlined, -1)
		case "purple bold":
			formatted = strings.Replace(formatted, "<"+t+">", PurpleUnderlined, -1)
		case "white bold":
			formatted = strings.Replace(formatted, "<"+t+">", WhiteUnderlined, -1)

		case "red bold underline":
			formatted = strings.Replace(formatted, "<"+t+">", RedBoldUnderlined, -1)
		case "red underline bold":
			formatted = strings.Replace(formatted, "<"+t+">", RedBoldUnderlined, -1)
		case "green bold underline":
			formatted = strings.Replace(formatted, "<"+t+">", GreenBoldUnderlined, -1)
		case "green underline bold":
			formatted = strings.Replace(formatted, "<"+t+">", GreenBoldUnderlined, -1)
		case "yellow bold underline":
			formatted = strings.Replace(formatted, "<"+t+">", YellowBoldUnderlined, -1)
		case "yellow underline bold":
			formatted = strings.Replace(formatted, "<"+t+">", YellowBoldUnderlined, -1)
		case "cyan bold underline":
			formatted = strings.Replace(formatted, "<"+t+">", CyanBoldUnderlined, -1)
		case "cyan underline bold":
			formatted = strings.Replace(formatted, "<"+t+">", CyanBoldUnderlined, -1)
		case "blue bold underline":
			formatted = strings.Replace(formatted, "<"+t+">", BlueBoldUnderlined, -1)
		case "blue underline bold":
			formatted = strings.Replace(formatted, "<"+t+">", BlueBoldUnderlined, -1)
		case "purple bold underline":
			formatted = strings.Replace(formatted, "<"+t+">", PurpleBoldUnderlined, -1)
		case "purple underline bold":
			formatted = strings.Replace(formatted, "<"+t+">", PurpleBoldUnderlined, -1)
		case "white bold underline":
			formatted = strings.Replace(formatted, "<"+t+">", WhiteBoldUnderlined, -1)
		case "white underline bold":
			formatted = strings.Replace(formatted, "<"+t+">", WhiteBoldUnderlined, -1)
		}
	}
	reset := regexp.MustCompile(`</(blue|green|red|yellow|error|warning|success|bold|undeline)>`)
	formatted = reset.ReplaceAllString(formatted, Reset)
	return formatted
}
