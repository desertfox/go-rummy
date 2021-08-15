package plugins

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	zshPath = ".oh-my-zsh"
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

	cmd := exec.Command("/bin/sh", "install.sh")

	err := cmd.Start()
	Check(err)

	err = cmd.Wait()
	Check(err)
}

func (p *ZshPlugin) installZshrc() {
	p.MoveFiles()
}
