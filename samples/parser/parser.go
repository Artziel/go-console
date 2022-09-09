package main

import (
	"fmt"
	"strings"

	Console "github.com/artziel/go-console"
)

const htmlFile = `
Prefix <text green>A list of colours</text> sufix
`

func main() {

	data := Console.Parse(htmlFile)
	parsed := strings.Join(data, "")
	fmt.Println(parsed)
}
