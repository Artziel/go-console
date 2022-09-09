package GoConsole

import (
	"strings"

	"golang.org/x/net/html"
)

func Parse(text string) (data []string) {
	var vals []string

	tkn := html.NewTokenizer(strings.NewReader(text))

	var colored bool

	for {

		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return vals

		case tt == html.StartTagToken:

			t := tkn.Token()
			colored = t.Data == "text"

		case tt == html.TextToken:

			t := tkn.Token()

			if colored {
				vals = append(vals, "<COLORED>"+t.Data+"</COLORED>")
			} else {
				vals = append(vals, t.Data)
			}

			colored = false
		}
	}
}
