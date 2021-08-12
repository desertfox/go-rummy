package rummy

import (
	p "github.com/go-rummy/pkg/plugins"
)

type Installer interface {
	Install()
}

func Go(sourceDir string, plugins []string) {

	Plugins := []Installer{p.NewBashPlugin(sourceDir), p.NewVimPlugin(sourceDir)}

	installPlugins(Plugins)
}

func installPlugins(plugins []Installer) {
	for _, plugin := range plugins {
		plugin.Install()
	}
}
