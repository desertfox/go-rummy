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

func NewVimPlugin(sourceDir string, destDir string) types.Installer {

	plugin := &types.PluginData{
		Name: "vim",
		SourceFilesDir: sourceDir,
		DestFilesDir: destDir,
	}

	vp := &VimPlugin{plugin}

	vp.AddFileToMove( ".vimrc", ".vimrc", false )

	return vp
}

func (p VimPlugin) Install() {
	fmt.Printf("Install: %v\n", p)

	p.installVimrc()
	p.installVimPlug()
}

func (p VimPlugin) installVimrc() {
	for _, ftm := range p.Files {
		Move(ftm)
	}
}

func (p VimPlugin) installVimPlug() {
	vimPlugPath := filepath.Join(p.DestFilesDir, vimPlugDestFile)
	
	if _, err := os.Stat(vimPlugPath); err == nil {
		fmt.Printf("%v file already exists\n", vimPlugPath)
		return
	}

	DownloadFile(vimPlugUrl, vimPlugPath)
}
