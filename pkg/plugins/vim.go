package plugins

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

	fmt.Printf("NewVim %v\n", p)

	return p
}

func (p VimPlugin) Install() {
	fmt.Printf("Install: %v\n", p)

	p.installVimrc()
	p.installVimPlug()
}

func (p VimPlugin) installVimrc() {
	fmt.Printf("installVimrc: %v\n", p)

	for _, file := range p.Data.FileNames {

		sourceFile := filepath.Join(p.Config.Cwd, p.Config.SourceFilesDir, file)
		destFile := filepath.Join(p.Dest, file)

		Move(sourceFile, destFile, p.Data.Overwrite)
	}
}

func (p VimPlugin) installVimPlug() {
	var VimAutoloadPath = filepath.Join(p.Dest, vimPlugDestFile)

	if _, err := os.Stat(VimAutoloadPath); os.IsNotExist(err) {
		fmt.Printf("vim plug.vim file already exists. bailing\n")
		return
	}

	resp, err := http.Get(vimPlugUrl)
	Check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	Check(err)
	err = ioutil.WriteFile(VimAutoloadPath, body, 0644)
	Check(err)
}
