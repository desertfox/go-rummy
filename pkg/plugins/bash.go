package plugins

import (
	"fmt"
	"os"
)

type BashPlugin struct {
	fileNames []string
	overwrite bool
}

func NewBashPlugin() *BashPlugin {
	bashPlugin := &BashPlugin{}

	bashPlugin.SetPluginFiles()

	fmt.Printf("NewBash %v\n", bashPlugin)

	return bashPlugin
}

func (b BashPlugin) GetPluginName() string {
	return "bash"
}

func (b *BashPlugin) SetPluginFiles() []string {

	b.fileNames = []string{".bash_aliases"}

	fmt.Printf("SetPluginFiles: %v\n", b)

	return b.fileNames
}

func (b *BashPlugin) Install() {
	fmt.Printf("Install: %v\n", b)
	b.installBashAliases()
}

func (b *BashPlugin) installBashAliases() {
	fmt.Printf("installBashA: %v\n", b)

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

}
