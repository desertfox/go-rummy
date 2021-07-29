package plugins

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	vimPlugUrl      = "https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim"
	vimPlugDestFile = ".vim/autoload/plug.vim"
)

type VimPlugin struct {
	fileNames []string
	overwrite bool
}

func NewVimPlugin() *VimPlugin {
	v := &VimPlugin{}

	v.SetPluginFiles()

	fmt.Printf("NewVim %v\n", v)

	return v
}

func (v *VimPlugin) GetPluginName() string {
	return "vim"
}

func (v *VimPlugin) SetPluginFiles() []string {

	v.fileNames = []string{".vimrc"}

	fmt.Printf("SetPluginFiles: %v\n", v)

	return v.fileNames
}

func (v *VimPlugin) Install() {
	fmt.Printf("Install: %v\n", v)

	v.installVimrc()
	v.installVimPlug()
}

func (v *VimPlugin) installVimrc() {

	fmt.Printf("installVimrc: %v\n", v)

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
}

func (v *VimPlugin) installVimPlug() {
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
