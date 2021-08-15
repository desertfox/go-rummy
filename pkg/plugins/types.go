package plugins

type PluginData struct {
	Name string
	FilesArray
	SourceFilesDir string
	DestFilesDir   string
}

type FilesArray struct {
	Files []FileToMove
}

type FileToMove struct {
	From      string
	To        string
	Overwrite bool
}

type FileToDownload struct {
	Url       string
	To        string
	Overwrite bool
}

type Installer interface {
	Install()
}
