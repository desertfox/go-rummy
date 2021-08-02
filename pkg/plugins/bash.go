package plugins

import (
	"fmt"
	//	"os"

	"github.com/go-rummy/pkg/types"
)

type BashPlugin struct {
	Data types.PluginData
}

func NewBashPlugin() *BashPlugin {
	p := &BashPlugin{
		Data: types.PluginData{
			Name:      "bash",
			FileNames: []string{".bash_aliases"},
			Overwrite: false,
		},
	}

	fmt.Printf("NewBash %v\n", p)

	return p
}

func (p BashPlugin) Install() {
	fmt.Printf("Install: %v\n", p)
	p.installBashAliases()
}

func (p BashPlugin) installBashAliases() {
	fmt.Printf("installBashA: %v\n", p)
	/*
		for _, file := range b.fileNames {
			mp := &PluginMove{
				sourcedir:  "dot-files",
				sourcefile: file,
				destdir:    os.Getenv("HOME"),
				dest:       file,
				overwrite:  b.overwrite,
			}

			mp.Move()
		}
	*/

}
