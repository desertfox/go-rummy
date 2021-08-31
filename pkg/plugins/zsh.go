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

func NewZshPlugin(destDir string, overwrite bool) Installer {
	p := &ZshPlugin{
		&PluginData{
			Name:         "zsh",
			DestFilesDir: destDir,
		},
	}

	p.AddConfigToCreate(&zshrc, ".zshrc", overwrite)
	p.AddConfigToCreate(&p10kzsh, ".p10k.zsh", overwrite)

	return p
}

func (p *ZshPlugin) Install() {
	err := p.installOhMyZSH()
	if err != nil {
		return
	}

	p.installZshrc()
	p.installPowerline()
}

func (p *ZshPlugin) installOhMyZSH() error {
	fullPath := p.BuildDestWithFile(zshPath)
	if _, err := os.Stat(fullPath); err == nil {
		fmt.Printf("%v file already exists\n", fullPath)
		return nil
	}

	f, err := ioutil.TempFile("", "tmp")
	Check(err)
	defer f.Close()

	installByte := DownloadFile("https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh")

	_, err = f.Write(installByte)
	Check(err)

	cmd := exec.Command("/bin/sh", f.Name())

	err = cmd.Start()
	Check(err)

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (p *ZshPlugin) installZshrc() {
	p.CreateConfigs()
}

func (p *ZshPlugin) installPowerline() {
	fullPath := p.BuildDestWithFile(zshPath)
	fullPath = filepath.Join(fullPath, "/custom/themes/powerlevel10k")

	if _, err := os.Stat(fullPath); err == nil {
		fmt.Printf("%v file already exists\n", fullPath)
		return
	}

	cmd := exec.Command("/usr/bin/git", "clone", "--depth=1", powerlineurl, fullPath)

	err := cmd.Start()
	Check(err)

	err = cmd.Wait()
	Check(err)

}
