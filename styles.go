package verbose

// Success is a success style
var Success = Style{
	Foreground: 2,
	Background: 0,
	Text:       "Success",
	Level:      4,
}

// Info is a warning style
var Info = Style{
	Foreground: 4,
	Background: 0,
	Text:       "Info",
	Level:      3,
}

// Warn is a warning style
var Warn = Style{
	Foreground: 3,
	Background: 0,
	Text:       "Warning",
	Level:      2,
}

// Error is an error style
var Error = Style{
	Foreground: 7,
	Background: 1,
	Text:       "Error",
	Level:      1,
}
