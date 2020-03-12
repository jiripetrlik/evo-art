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
				<td>
					<a href="?chromosome={{index .Chromosomes 0}}">
						<img src="image.png?chromosome={{index .Chromosomes 0}}">
					</a>
				</td>
				<td>
					<a href="?chromosome={{index .Chromosomes 1}}">
						<img src="image.png?chromosome={{index .Chromosomes 1}}">
					</a>
				</td>
			</tr>
			<tr>
				<td>
					<a href="?chromosome={{index .Chromosomes 2}}">
						<img src="image.png?chromosome={{index .Chromosomes 2}}">
					</a>
				</td>
				<td>
					<a href="?chromosome={{index .Chromosomes 3}}">
						<img src="image.png?chromosome={{index .Chromosomes 3}}">
					</a>
				</td>
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
