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
	powerlineurl = "git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/themes/powerlevel10k"
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

func NewZshPlugin(sourceDir string, destDir string, overwrite bool) Installer {
	sourceDir = filepath.Join(sourceDir, "zsh")

	plugin := &PluginData{
		Name:           "zsh",
		SourceFilesDir: sourceDir,
		DestFilesDir:   destDir,
	}

	zp := &ZshPlugin{plugin}

	zp.AddConfigToCreate(&zshrc, ".zshrc", overwrite)
	zp.AddConfigToCreate(&p10kzsh, ".p10k.zsh", overwrite)

	return zp
}

func (p *ZshPlugin) Install() {
	p.installOhMyZSH()
	p.installZshrc()
}

func (p *ZshPlugin) installOhMyZSH() {
	fullPath := filepath.Join(os.Getenv("HOME"), zshPath)
	if _, err := os.Stat(fullPath); err == nil {
		fmt.Printf("%v file already exists\n", fullPath)
		return
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
	Check(err)
}

func (p *ZshPlugin) installZshrc() {
	p.CreateConfigs()
}
