package main

var ResultsTemplate = `

<!DOCTYPE html>
<html lang="en">
<head>
		<title>Campwiz2</title>
		<meta charset="utf-8">
</head>
<body>
		{{range .}}
		------------------------------------------------------------------------------------
				{{.Name}} 
		------------------------------------------------------------------------------------
		{{.BookRecord.SRating}}

				<br>
		{{end}}
</body>
</html>
`
