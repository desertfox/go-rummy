package plugins

import (
	"fmt"
	"os"

	"github.com/go-rummy/pkg/types"
)

type BashPlugin struct {
	*types.PluginData
}

func NewBashPlugin(sourceDir string) types.Installer {

	plugin := types.NewPlugin("bash", []string{".bash_aliases"}, false, os.Getenv("HOME"), sourceDir)

	return &BashPlugin{plugin}
}

func (p *BashPlugin) Install() {
	fmt.Printf("Install: %v\n", p)

	p.installBashAliases()
}

func (p *BashPlugin) installBashAliases() {
	for _, file := range p.FileNames {
		sourceFile := p.BuildSourceWithFile(file)
		destFile := p.BuildDestWithFile(file)

		Move(sourceFile, destFile, p.Overwrite)
	}
}
