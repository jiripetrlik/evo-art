package gui

const choicePageTemplate string = `
<!doctype html>
<html>
	<head>
		<title>EVO ART</title>
		<link rel="stylesheet" type="text/css" href="style.css">
	</head>
	<body>
		<h1>EVO ART</h1>
		<table class="choice_table">
			<tr>
				<td><img src="image.png?chromosome={{index .Chromosomes 1 }}"></td>
				<td><img src="image.png?chromosome={{index .Chromosomes 1}}"></td>
			</tr>
			<tr>
				<td><img src="image.png?chromosome={{index .Chromosomes 2}}"></td>
				<td><img src="image.png?chromosome={{index .Chromosomes 3}}"></td>
			</tr>
		</table>
	</body>
</html>
`

const css = `
h1 {
	text-align: center;
}

table.choice_table {
	margin-right: auto;
	margin-left: auto;
}
`
