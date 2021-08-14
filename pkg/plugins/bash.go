package plugins

import (
	"fmt"

	"github.com/go-rummy/pkg/types"
)

type BashPlugin struct {
	*types.PluginData
}

func NewBashPlugin(sourceDir string, destDir string) types.Installer {
	plugin := &types.PluginData{
		Name:           "bash",
		SourceFilesDir: sourceDir,
		DestFilesDir:   destDir,
	}

	bp := &BashPlugin{plugin}

	bp.AddFileToMove(".bash_aliases", ".bash_aliases", false)

	return bp
}

func (p *BashPlugin) Install() {
	fmt.Printf("Install: %v\n", p)

	p.installBashAliases()
}

func (p *BashPlugin) installBashAliases() {
	for _, ftm := range p.Files {
		Move(ftm)
	}
}
