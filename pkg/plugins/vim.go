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
	Data   types.PluginData
	Dest   string
	Config types.Config
}

func NewVimPlugin(config types.Config) *VimPlugin {
	p := &VimPlugin{
		Data: types.PluginData{
			Name:      "vim",
			FileNames: []string{".vimrc"},
			Overwrite: false,
		},
		Dest:   os.Getenv("HOME"),
		Config: config,
	}

	return p
}

func (p VimPlugin) Install() {
	fmt.Printf("Install: %v\n", p)

	p.installVimrc()
	p.installVimPlug()
}

func (p VimPlugin) installVimrc() {
	for _, file := range p.Data.FileNames {
		sourceFile := filepath.Join(p.Config.Cwd, p.Config.SourceFilesDir, file)
		destFile := filepath.Join(p.Dest, file)

		Move(sourceFile, destFile, p.Data.Overwrite)
	}
}

func (p VimPlugin) installVimPlug() {
	vimPlugPath := filepath.Join(p.Dest, vimPlugDestFile)

	DownloadFile(vimPlugUrl, vimPlugPath)
}
