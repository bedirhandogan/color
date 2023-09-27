package color

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os"
)

// Support ANSI for color code in Windows terminal.
func init() {
	var originalMode uint32
	stdout := windows.Handle(os.Stdout.Fd())

	defer func(console windows.Handle, mode uint32) {
		err := windows.SetConsoleMode(console, mode)
		if err != nil {
			fmt.Println(err)
		}
	}(stdout, originalMode)

	err := windows.GetConsoleMode(stdout, &originalMode)

	if err != nil {
		fmt.Println(err)
	}

	err = windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)

	if err != nil {
		fmt.Println(err)
	}
}
