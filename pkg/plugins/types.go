package plugins

type PluginData struct {
	Name           string
	Files          []FileToMove
	SourceFilesDir string
	DestFilesDir   string
}

type FileToMove struct {
	From      string
	To        string
	Overwrite bool
}

type Installer interface {
	Install()
}
