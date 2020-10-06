package source

import (
	"fmt"
)

func GetDescription() string {
	return fmt.Sprintf("%s%s", newMsg, `

My name is {{.name}}, and I am {{.age}} years old. I like:
{{define "interestsList"}}- {{.}}{{end}}{{range $i := .interests}}
    {{template "interestsList" $i}}
{{end}}
This is my favorite quote: "{{.quote}}"
`)
}