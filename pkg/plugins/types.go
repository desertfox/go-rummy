package plugins

type PluginData struct {
	Name    string
	Configs []ConfigToCreate
}

type ConfigToCreate struct {
	Source    *string
	To        string
	Overwrite bool
}

type Installer interface {
	Install(destDir string, overwrite bool) error
}
