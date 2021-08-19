package plugins

type PluginData struct {
	Name           string
	Configs        []ConfigToCreate
	SourceFilesDir string
	DestFilesDir   string
}

type ConfigToCreate struct {
	Source    *string
	To        string
	Overwrite bool
}

type Installer interface {
	Install()
}
