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

func NewVimPlugin() Installer {
	return &VimPlugin{&PluginData{
		Name: "vim",
	}}
}

func (p *VimPlugin) Install(destDir string, overwrite bool) error {
	p.AddConfigToCreate(&vimrc, p.buildDestPath(destDir, ".vimrc"), overwrite)

	if err := p.installVimrc(); err != nil {
		return err
	}

	return p.installVimPlug(destDir)
}

func (p VimPlugin) installVimrc() error {
	return p.CreateConfigs()
}

func (p VimPlugin) installVimPlug(destDir string) error {
	vimPlugPath := p.buildDestPath(destDir, vimPlugDestFile)

	if _, err := os.Stat(vimPlugPath); err == nil {
		fmt.Printf("%v file already exists\n", vimPlugPath)
		return nil
	}

	pathSlice := strings.Split(vimPlugPath, "/")
	path := strings.Join(pathSlice[0:len(pathSlice)-1], "/")

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	installByte, err := DownloadFile(vimPlugUrl)
	if err != nil {
		return err
	}

	f, err := os.Create(vimPlugPath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(installByte)
	return err
}
