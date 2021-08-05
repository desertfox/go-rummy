package rummy

import (
	"github.com/go-rummy/pkg/plugins"
	"github.com/go-rummy/pkg/types"
)

func GoRummy(wd string) {

	config := types.NewConfig(wd)

	plugins := []types.Installer{plugins.NewBashPlugin(*config), plugins.NewVimPlugin(*config)}

	installPlugins(plugins)
}

func installPlugins(plugins []types.Installer) {
	for _, plugin := range plugins {
		plugin.Install()
	}
}
