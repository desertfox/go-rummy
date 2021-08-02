package plugins

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/go-rummy/pkg/types"
)

var (
	vimPlugUrl      = "https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim"
	vimPlugDestFile = ".vim/autoload/plug.vim"
)

type VimPlugin struct {
	Data types.PluginData
}

func NewVimPlugin() *VimPlugin {
	p := &VimPlugin{
		Data: types.PluginData{
			Name:      "vim",
			FileNames: []string{".vimc"},
			Overwrite: false,
		},
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
	/*
		for _, file := range v.fileNames {
			mp := &PluginMove{
				sourcedir:  "dot-files",
				sourcefile: file,
				destdir:    os.Getenv("HOME"),
				dest:       file,
				overwrite:  v.overwrite,
			}

			mp.Move()
		}
	*/
}

func (p VimPlugin) installVimPlug() {
	var VimAutoloadPath = strings.Join([]string{os.Getenv("HOME"), vimPlugDestFile}, "/")

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
