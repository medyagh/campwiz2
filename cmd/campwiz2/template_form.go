package main

var searchFormTemplate = `

<!DOCTYPE html>
<html lang="en">
<head>
		<title>Campwiz2</title>
		<meta charset="utf-8">
</head>
<body>
<h1> Welcome to Campwiz2 </h1>
<h3>Lets Go Camping </h3>

<form action="/" method="post">
    <fieldset>
        <label for="dates">Dates</label>
        <input type="date" id="dates" name="dates" value="{{ .Today }}" min="{{ .Today }}">
    </fieldset>

    <fieldset>
        <label>Nights</label>
        <input type="number" name="nights" min="1" max="7" step="1" value="1" />
    </fieldset>

    <fieldset>
        <label for="distance">Distance</label>
        <select name="distance" id="distance">
            <option value="50">50 miles</option>
            <option value="100">100 miles</option>
            <option value="150">150 miles</option>
			<option value="250" selected >250 miles</option>
			<option value="350" >350 miles</option>
        </select>
    </fieldset>

    <fieldset>
        <label>Type</label>
        <input type="checkbox" name="standard" checked >Standard</input>
        <input type="checkbox" name="group" >Group</input>
        <input type="checkbox" name="walk-in" >Walk-in</input>
        <input type="checkbox" name="boat-in" >Boat-in</input>
        </label>
    </fieldset>

    <button type="submit">Search</button>
</form>


</body>
</html>
`
