package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/c-bata/go-prompt"
)

func templating(input string) string {
	if !strings.Contains(input, "{{") {
		input = "{{" + input + "}}"
	}

	tpl, err := template.New("REPL").Funcs(sprig.FuncMap()).Parse(input)
	if err != nil {
		fmt.Println("ERROR", err)
		return ""
	}

	var b bytes.Buffer
	err = tpl.Execute(&b, nil)
	if err != nil {
		fmt.Println("ERROR", err)
		return ""
	}

	return b.String() + "\n"
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			fmt.Print(templating(arg))
		}
	} else {
		fmt.Println("Use CTRL + D to exit")
		p := prompt.New(
			func(in string) {
				switch strings.ToLower(in) {
				case "/help":
					for _, cmd := range suggest {
						fmt.Printf("%s - %s\n", cmd.Text, cmd.Description)
					}
				case "/quit":
					os.Exit(0)
				default:
					if in != "" {
						fmt.Print(templating(in))
					}
				}
			},
			func(d prompt.Document) []prompt.Suggest {
				return prompt.FilterHasPrefix(suggest, d.GetWordBeforeCursor(), false)
			},
			prompt.OptionPrefix("> "),
			prompt.OptionTitle("Sprig REPL"),
		)
		p.Run()
	}
}
