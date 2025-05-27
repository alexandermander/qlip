// formatText()
//
//butten myButton

document.getElementById("myButton").addEventListener("click", function() {
	// get the text in the "myText" id
	
	const documentText = document.getElementById("myText").value;

	const lines = documentText.split("\n");
	let formattedLines = [];
	for (let i = 0; i < lines.length; i++) {
		// remove leading and trailing whitespace
		lines[i] = lines[i].trim();
		formattedLines.push(lines[i]);
	}
	// join the lines with a newline character
	const formattedText = formattedLines.join(" ");
	console.log(formattedText);
	//split byt "- "
	const splitLines = formattedText.split("- ");
	// remove empty split
	const nonEmptyLines = splitLines.filter(line => line.trim() !== "");
	const joinedText = nonEmptyLines.join("");

	// set the text in the "myText" ge
	document.getElementById("myText").value = joinedText;

});
