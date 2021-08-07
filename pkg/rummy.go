package rummy

import (
	"github.com/go-rummy/pkg/plugins"
	"github.com/go-rummy/pkg/types"
)

func Go(wd string, dotFiles string) {
	config := types.NewConfig(wd, dotFiles)

	plugins := []types.Installer{plugins.NewBashPlugin(*config), plugins.NewVimPlugin(*config)}

	installPlugins(plugins)
}

func installPlugins(plugins []types.Installer) {
	for _, plugin := range plugins {
		plugin.Install()
	}
}
