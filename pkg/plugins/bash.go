package plugins

import _ "embed"

var (
	//go:embed dot-files/bash/.bash_aliases
	bashAliases string
)

type BashPlugin struct {
	*PluginData
}

func NewBashPlugin() Installer {
	return &BashPlugin{&PluginData{
		Name: "bash",
	}}
}

func (p *BashPlugin) Install(destDir string, overwrite bool) error {
	p.AddConfigToCreate(&bashAliases, p.buildDestPath(destDir, ".bash_aliases"), overwrite)

	return p.installBashAliases()
}

func (p *BashPlugin) installBashAliases() error {
	return p.CreateConfigs()
}
