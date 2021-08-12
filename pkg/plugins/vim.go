package plugins

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-rummy/pkg/types"
)

var (
	vimPlugUrl      = "https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim"
	vimPlugDestFile = ".vim/autoload/plug.vim"
)

type VimPlugin struct {
	*types.PluginData
}

func NewVimPlugin(sourceDir string) types.Installer {

	plugin := types.NewPlugin("vim", []string{".vimrc"}, false, os.Getenv("HOME"), sourceDir)

	return &VimPlugin{plugin}
}

func (p VimPlugin) Install() {
	fmt.Printf("Install: %v\n", p)

	p.installVimrc()
	p.installVimPlug()
}

func (p VimPlugin) installVimrc() {
	for _, file := range p.FileNames {
		sourceFile := filepath.Join(p.SourceFilesDir, file)
		destFile := filepath.Join(p.Dest, file)

		Move(sourceFile, destFile, p.Overwrite)
	}
}

func (p VimPlugin) installVimPlug() {
	vimPlugPath := filepath.Join(p.Dest, vimPlugDestFile)

	DownloadFile(vimPlugUrl, vimPlugPath)
}
