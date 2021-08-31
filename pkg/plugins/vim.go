package plugins

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
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

func NewVimPlugin(destDir string, overwrite bool) Installer {

	plugin := &PluginData{
		Name:         "vim",
		DestFilesDir: destDir,
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
	vimPlugPath := p.BuildDestWithFile(vimPlugDestFile)

	if _, err := os.Stat(vimPlugPath); err == nil {
		fmt.Printf("%v file already exists\n", vimPlugPath)
		return
	}

	pathSlice := strings.Split(vimPlugPath, "/")
	path := strings.Join(pathSlice[0:len(pathSlice)-1], "/")

	err := os.MkdirAll(path, os.ModePerm)
	Check(err)

	installByte := DownloadFile(vimPlugUrl)

	f, err := os.Create(vimPlugPath)
	Check(err)
	defer f.Close()

	_, err = f.Write(installByte)
	Check(err)
}
