package main

var ResultsTemplate = `
{{range $index, $element := .Records}}
	{{$index}} 	{{.Name}}
{{end}}
`
