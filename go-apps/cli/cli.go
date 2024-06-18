package cli

import (
	"fmt"
	"os"

	"phpguru.net/go-apps/appbase"
	"phpguru.net/go-apps/help"
	"phpguru.net/go-apps/leapyear"
)

const (
	Help             = "help"
	Version          = "version"
	LeapYear         = "leapyear"
	BMI              = "bmi"
	CompoundInterest = "ci"
)

func getApps() map[string]appbase.App {
	apps := make(map[string]appbase.App)
	apps[Help] = help.NewHelpApp()
	// apps[Version] = "display version"
	apps[LeapYear] = leapyear.NewLeapApp()
	return apps
}

func getCommands(apps map[string]appbase.App) map[string]string {
	commands := make(map[string]string)
	for appName, appInstance := range apps {
		commands[appName] = appInstance.GetAppName()
	}
	return commands
}

func ReadArgs() {
	args := os.Args
	apps := getApps()
	commands := getCommands(apps)
	displayHelp := func() {
		apps[Help].Run(commands)
	}

	if len(args) < 2 {
		displayHelp()
		os.Exit(1)
	}
	// otherwise second args ( index = 1 ) is the command
	switch args[1] {
	case LeapYear:
		apps[LeapYear].Run(help.ParseAny(os.Args[2:])...)
	case BMI:
		fmt.Println(args[1])
	case CompoundInterest:
		fmt.Println(args[1])
	default:
		displayHelp()
	}
}
