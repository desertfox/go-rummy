package rummy

import (
	"errors"

	p "github.com/go-rummy/pkg/plugins"
)

type Installer interface {
	Install(destDir string, overwrite bool) error
}

type Rummy struct {
	InstallList []string
	DestDir     string
	Overwrite   bool
}

func (r *Rummy) Go() error {
	plugins, err := r.selectPluginFromInstallList()
	if err != nil {
		return err
	}

	for _, plugin := range *plugins {
		err := plugin.Install(r.DestDir, r.Overwrite)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Rummy) selectPluginFromInstallList() (*map[string]Installer, error) {
	list := make(map[string]Installer)

	list["git"] = p.NewGitPlugin()
	list["vim"] = p.NewVimPlugin()
	list["zsh"] = p.NewZshPlugin()
	list["bash"] = p.NewBashPlugin()

	if r.InstallList[0] == "all" {
		return &list, nil
	}

	sublist := make(map[string]Installer, len(r.InstallList))
	for _, v := range r.InstallList {
		if _, exists := list[v]; exists {
			sublist[v] = list[v]
		} else {
			return nil, errors.New("No plugin found for " + v)
		}
	}
	return &sublist, nil
}
