package help

import (
	"fmt"
	"strings"

	"phpguru.net/go-apps/appbase"
)

type Help struct {
	appbase.AppInformation
}

// Run implements appbase.App.
func (h *Help) Run(args ...any) {
	if len(args) == 0 {
		fmt.Println("Error: No arguments provided to Run")
		return
	}
	commandMap, ok := args[0].(map[string]string)
	if !ok {
		fmt.Println("Error: First argument is not a map[string]string")
		return
	}

	space := strings.Repeat(" ", 4)
	newLine := "\n"
	output := "Go base is tool includes useful applications." + newLine
	output += "Usage:" + newLine
	output += space + "gobase <command> [arguments]" + newLine
	output += "The commands are:" + newLine
	maxCommandLength := 0
	for command := range commandMap {
		l := len(command)
		if l > maxCommandLength {
			maxCommandLength = l
		}
	}
	paddingSpace := strings.Repeat(" ", maxCommandLength) + space
	for command, decription := range commandMap {
		output += space + command + paddingSpace + decription + newLine
	}
	fmt.Println(output)
}

func NewHelpApp() appbase.App {
	return &Help{
		appbase.AppInformation{
			Name: "display help message",
		},
	}
}
