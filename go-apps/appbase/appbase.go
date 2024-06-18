package appbase

type AppInformation struct {
	Name string
}

func (a *AppInformation) GetAppName() string {
	return a.Name
}

type App interface {
	GetAppName() string
	Run(args ...any)
}
