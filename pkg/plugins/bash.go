package plugins

type BashPlugin struct {
	*PluginData
}

func NewBashPlugin(sourceDir string, destDir string, overwrite bool) Installer {
	plugin := &PluginData{
		Name:           "bash",
		SourceFilesDir: sourceDir,
		DestFilesDir:   destDir,
	}

	bp := &BashPlugin{plugin}

	bp.AddFileToMove(".bash_aliases", ".bash_aliases", overwrite)

	return bp
}

func (p *BashPlugin) Install() {
	p.installBashAliases()
}

func (p *BashPlugin) installBashAliases() {
	p.MoveFiles()
}
