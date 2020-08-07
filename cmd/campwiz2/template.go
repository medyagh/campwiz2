package main

var ResultsTemplate = `
  
<!DOCTYPE html>
<html lang="en">
<head>
		<title>Campwiz2</title>
		<meta charset="utf-8">
</head>
<body>

<table  id="grid-basic" class="table table-condensed table-hover table-striped">
    <thead>
        <tr>
            <th data-column-id="id" data-type="numeric">ID</th>
            <th data-column-id="name">Name</th>
			<th data-column-id="Proximity" data-order="desc">Proximity</th>
			<th data-column-id="Availability">Availability</th>
			<th data-column-id="VerifiableAvailability">VerifiableAvailability</th>
        </tr>
    </thead>
    <tbody>
	{{range $index, $element := .}}
	<tr>
		<td>{{$index}}</td>
		<td>{{.Name}}</td>
		<td>{{.Proximity}}</td>
		<td>{{.Details.Availability}}</td>
		<td>{{.Details.VerifiableAvailability}}</td>
	</tr>	    
	{{end}}
    </tbody>
</table>

</body>
</html>
`
