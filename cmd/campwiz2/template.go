package main

var ResultsTemplate = `
{{range $index, $element := .}}
	{{$index}} 	{{.Name}}
{{end}}
`
