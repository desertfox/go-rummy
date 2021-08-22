package plugins

import _ "embed"

var (
	//go:embed dot-files/bash/.bash_aliases
	bashAliases string
)

type BashPlugin struct {
	*PluginData
}

func NewBashPlugin(destDir string, overwrite bool) Installer {
	plugin := &PluginData{
		Name:         "bash",
		DestFilesDir: destDir,
	}

	bp := &BashPlugin{plugin}

	bp.AddConfigToCreate(&bashAliases, ".bash_aliases", overwrite)

	return bp
}

func (p *BashPlugin) Install() {
	p.installBashAliases()
}

func (p *BashPlugin) installBashAliases() {
	p.CreateConfigs()
}
