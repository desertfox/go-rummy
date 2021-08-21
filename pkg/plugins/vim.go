package plugins

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
)

var (
	vimPlugUrl      = "https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim"
	vimPlugDestFile = ".vim/autoload/plug.vim"
)

var (
	//go:embed dot-files/vim/.vimrc
	vimrc string
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

	vp.AddConfigToCreate(&vimrc, ".vimrc", overwrite)

	return vp
}

func (p VimPlugin) Install() {
	p.installVimrc()
	p.installVimPlug()
}

func (p VimPlugin) installVimrc() {
	p.CreateConfigs()
}

func (p VimPlugin) installVimPlug() {
	vimPlugPath := filepath.Join(p.DestFilesDir, vimPlugDestFile)

	if _, err := os.Stat(vimPlugPath); err == nil {
		fmt.Printf("%v file already exists\n", vimPlugPath)
		return
	}

	//DownloadFile(vimPlugUrl, vimPlugPath)
}
