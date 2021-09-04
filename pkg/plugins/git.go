package plugins

import _ "embed"

var (
	//go:embed dot-files/git/.gitconfig
	gitconfig string
)

type GitPlugin struct {
	*PluginData
}

func NewGitPlugin() Installer {
	return &GitPlugin{&PluginData{
		Name: "git",
	}}
}

func (p *GitPlugin) Install(destDir string, overwrite bool) error {
	p.AddConfigToCreate(&gitconfig, p.buildDestPath(destDir, ".gitconfig"), overwrite)

	return p.CreateConfigs()
}
