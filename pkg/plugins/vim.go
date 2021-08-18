package plugins

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	vimPlugUrl      = "https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim"
	vimPlugDestFile = ".vim/autoload/plug.vim"
)

type VimPlugin struct {
	*PluginData
}

func NewVimPlugin(sourceDir string, destDir string, overwrite bool) Installer {

	plugin := &PluginData{
		Name:           "vim",
		SourceFilesDir: sourceDir,
		DestFilesDir:   destDir,
	}

	vp := &VimPlugin{plugin}

	vp.AddFileToMove(".vimrc", ".vimrc", overwrite)

	return vp
}

func (p VimPlugin) Install() {
	fmt.Printf("Install: %v\n", p)

	p.installVimrc()
	p.installVimPlug()
}

func (p VimPlugin) installVimrc() {
	p.MoveFiles()
}

func (p VimPlugin) installVimPlug() {
	vimPlugPath := filepath.Join(p.DestFilesDir, vimPlugDestFile)

	if _, err := os.Stat(vimPlugPath); err == nil {
		fmt.Printf("%v file already exists\n", vimPlugPath)
		return
	}

	DownloadFile(vimPlugUrl, vimPlugPath)
}
