package main

import (
	"fmt"

	"github.com/willbarkoff/verbose"
)

func main() {
	// First, let's just print some messages.

	verbose.Success.Println("This is a success message.")
	verbose.Info.Println("This is an informational message.")
	verbose.Warn.Println("This is a warning.")
	verbose.Error.Println("This is an error.")

	fmt.Println() //Just print a space.

	// Now, let's define a custom style

	xmas := verbose.Style{
		Foreground: 1,
		Background: 2,
		Text:       "Christmas",
		Level:      1,
	}

	xmas.Println("Merry Christmas!")

	fmt.Println() //Just print a space.

	// What if we dont' want to use colors?
	verbose.ColorsEnabled = false

	verbose.Success.Println("This is a success message.")
	verbose.Info.Println("This is an informational message.")
	verbose.Warn.Println("This is a warning.")
	verbose.Error.Println("This is an error.")

	verbose.ColorsEnabled = true

	fmt.Println() //Just print a space.

	// Maybe I don't want to print success and info messages unless the user explicitly enables it?

	verbose.MaximumLogLevel = 2

	verbose.Success.Println("This is a success message.")
	verbose.Info.Println("This is an informational message.")
	verbose.Warn.Println("This is a warning.")
	verbose.Error.Println("This is an error.")
}
