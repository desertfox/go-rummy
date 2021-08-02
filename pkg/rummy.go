package rummy

import (
	"os"
	//	"fmt"

	"github.com/go-rummy/pkg/plugins"
	"github.com/go-rummy/pkg/types"
)

func GoRummy() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	rb := types.NewRepoBase(path)

	plugins := []types.RummyPlugin{plugins.NewBashPlugin(), plugins.NewVimPlugin()}

	rc := rb.NewRummyConfig(plugins)

	installPlugins(rc)
}

func installPlugins(rc *types.RummyConfig) {
	for _, plugin := range rc.Plugins {
		plugin.Install()
	}
}
