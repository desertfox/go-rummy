package plugins

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-rummy/pkg/types"
)

type BashPlugin struct {
	Data   types.PluginData
	Dest   string
	Config types.Config
}

func NewBashPlugin(config types.Config) *BashPlugin {
	p := &BashPlugin{
		Data: types.PluginData{
			Name:      "bash",
			FileNames: []string{".bash_aliases"},
			Overwrite: false,
		},
		Dest:   os.Getenv("HOME"),
		Config: config,
	}

	return p
}

func (p *BashPlugin) Install() {
	fmt.Printf("Install: %v\n", p)

	p.installBashAliases()
}

func (p *BashPlugin) installBashAliases() {
	for _, file := range p.Data.FileNames {
		sourceFile := filepath.Join(p.Config.Cwd, p.Config.SourceFilesDir, file)
		destFile := filepath.Join(p.Dest, file)

		Move(sourceFile, destFile, p.Data.Overwrite)
	}
}
