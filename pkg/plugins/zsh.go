package plugins

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	zshPath      = ".oh-my-zsh"
	powerlineurl = "https://github.com/romkatv/powerlevel10k.git"
)

var (
	//go:embed dot-files/zsh/.zshrc
	zshrc string
	//go:embed dot-files/zsh/.p10k.zsh
	p10kzsh string
)

type ZshPlugin struct {
	*PluginData
}

func NewZshPlugin() Installer {
	return &ZshPlugin{&PluginData{
		Name: "zsh",
	}}
}

func (p *ZshPlugin) Install(destDir string, overwrite bool) error {
	p.AddConfigToCreate(&zshrc, p.buildDestPath(destDir, ".zshrc"), overwrite)
	p.AddConfigToCreate(&p10kzsh, p.buildDestPath(destDir, ".p10k.zsh"), overwrite)

	err := p.installOhMyZSH(destDir)
	if err != nil {
		return err
	}

	err = p.installZshrc()
	if err != nil {
		return err
	}

	return p.installPowerline(destDir)
}

func (p *ZshPlugin) installOhMyZSH(destDir string) error {
	fullPath := p.buildDestPath(destDir, zshPath)
	if _, err := os.Stat(fullPath); err == nil {
		fmt.Printf("%v file already exists\n", fullPath)
		return nil
	}

	f, err := ioutil.TempFile("", "tmp")
	if err != nil {
		return err
	}
	defer f.Close()

	installByte, err := DownloadFile("https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh")
	if err != nil {
		return err
	}

	_, err = f.Write(installByte)
	if err != nil {
		return err
	}

	cmd := exec.Command("/bin/sh", f.Name())

	err = cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}

func (p *ZshPlugin) installZshrc() error {
	return p.CreateConfigs()
}

func (p *ZshPlugin) installPowerline(destDir string) error {
	fullPath := p.buildDestPath(destDir, zshPath)
	fullPath = filepath.Join(fullPath, "/custom/themes/powerlevel10k")

	if _, err := os.Stat(fullPath); err == nil {
		fmt.Printf("%v file already exists\n", fullPath)
		return nil
	}

	cmd := exec.Command("/usr/bin/git", "clone", "--depth=1", powerlineurl, fullPath)

	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}
