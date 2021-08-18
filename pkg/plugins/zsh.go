package plugins

import (
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

type ZshPlugin struct {
	*PluginData
}

func NewZshPlugin(sourceDir string, destDir string) Installer {
	sourceDir = filepath.Join(sourceDir, "zsh")

	plugin := &PluginData{
		Name:           "zsh",
		SourceFilesDir: sourceDir,
		DestFilesDir:   destDir,
	}

	zp := &BashPlugin{plugin}

	zp.AddFileToMove(".zshrc", ".zshrc", false)
	zp.AddFileToMove(".p10k.zsh", ".p10k.zsh", false)

	return zp
}

func (p *ZshPlugin) Install() {
	p.installOhMyZSH()
	p.installZshrc()
}

func (p *ZshPlugin) installOhMyZSH() {
	if _, err := os.Stat(zshPath); err == nil {
		fmt.Printf("%v file already exists\n", zshPath)
		return
	}

	dir, err := ioutil.TempDir("dir", "prefix")
	Check(err)

	DownloadFile("https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh", dir)

	cmd := exec.Command("/bin/sh", filepath.Join(dir, "install.sh"))

	err = cmd.Start()
	Check(err)

	err = cmd.Wait()
	Check(err)

	defer os.RemoveAll(dir)
}

func (p *ZshPlugin) installZshrc() {
	p.MoveFiles()
}
