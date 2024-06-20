package cli

import (
	"os"

	"phpguru.net/go-apps/appbase"
	"phpguru.net/go-apps/bmi"
	"phpguru.net/go-apps/compound_interest"
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
	apps[BMI] = bmi.NewBmiApp()
	apps[CompoundInterest] = compound_interest.NewCompoundInterestApp()
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
	payload := help.ParseAny(os.Args[2:])
	switch args[1] {
	case LeapYear:
		apps[LeapYear].Run(payload...)
	case BMI:
		apps[BMI].Run(payload...)
	case CompoundInterest:
		apps[CompoundInterest].Run(payload...)
	default:
		displayHelp()
	}
}
