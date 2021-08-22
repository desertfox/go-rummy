package plugins

import _ "embed"

var (
	//go:embed dot-files/git/.gitconfig
	gitconfig string
)

type GitPlugin struct {
	*PluginData
}

func NewGitPlugin(sourceDir string, destDir string, overwrite bool) Installer {

	plugin := &PluginData{
		Name:           "git",
		SourceFilesDir: sourceDir,
		DestFilesDir:   destDir,
	}

	zp := &GitPlugin{plugin}

	zp.AddConfigToCreate(&gitconfig, ".gitconfig", overwrite)

	return zp
}

func (p *GitPlugin) Install() {
	p.CreateConfigs()
}
