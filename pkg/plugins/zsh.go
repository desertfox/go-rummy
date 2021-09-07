package plugins

import _ "embed"

var (
	//go:embed dot-files/zsh/.zshrc
	zshrc string
	//go:embed dot-files/zsh/.p10k.zsh
	p10kzsh string
)

type ZshPlugin struct {
	*PluginData
}

func NewZshPlugin() Installer {
	return &ZshPlugin{&PluginData{
		Name: "zsh",
	}}
}

func (p *ZshPlugin) Install(destDir string, overwrite bool) error {
	p.AddConfigToCreate(&zshrc, p.buildDestPath(destDir, ".zshrc"), overwrite)
	p.AddConfigToCreate(&p10kzsh, p.buildDestPath(destDir, ".p10k.zsh"), overwrite)

	return p.CreateConfigs()

}
