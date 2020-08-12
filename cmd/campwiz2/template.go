package main

var ResultsTemplate = `

<!DOCTYPE html>
<html lang="en">
<head>
		<title>Campwiz2</title>
		<meta charset="utf-8">
		<link rel="stylesheet" type="text/css" href="https://cdn.datatables.net/1.10.21/css/jquery.dataTables.css">
  
		<script src="https://code.jquery.com/jquery-3.5.1.js"></script>
		<script src="https://cdn.datatables.net/1.10.21/js/jquery.dataTables.min.js"></script>
		<script src="https://cdn.datatables.net/1.10.19/js/dataTables.bootstrap4.min.js"></script>
		<script src="https://cdn.datatables.net/buttons/1.6.2/js/dataTables.buttons.min.js"></script>

</head>
<body>

<table id="myTable" class="display">
    <thead>
        <tr>
            <th>ID</th>
			<th>Name</th>
			<th>Proximity</th>
			<td>Availability</td>
			<td>VerifiableAvailability</td>
			<th>Book Rating</th>
			<th>Book Desc</th>
        </tr>
    </thead>
    <tbody>
	{{range $index, $element := .}}
		<tr>
			<th>{{$index}}</th>
			<th>{{.Name}}</th>
			<th>{{.Proximity}}</th>
			<td>{{.Details.Availability.Available}}</td>
			<td>{{.Details.VerifiableAvailability}}</td>
			<th>{{.BookRecord.SRating}}</th>
			<th>{{.BookRecord.Desc}}</th>
		</tr>
	{{end}}
    </tbody>
</table>
	
<script>
    $('#myTable').DataTable({
		"pageLength": 50

	});	
</script>
</body>
</html>
`
