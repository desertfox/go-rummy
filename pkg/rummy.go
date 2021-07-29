package rummy

import (
	"os"

	"github.com/go-rummy/pkg/plugins"
)

var cwd string

type localConfig struct {
	Cwd, DotfilesName string
}

type RummyPlugin interface {
	Install()
}

type RummyConfig struct {
	plugins []RummyPlugin
	base    *localConfig
}

func init() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cwd = path
}

func initLocalConfig() *localConfig {
	config := &localConfig{
		Cwd:          cwd,
		DotfilesName: "dot-files",
	}
	return config
}

func GoRummy() {

	plugins := []RummyPlugin{plugins.NewBashPlugin(), plugins.NewVimPlugin()}

	app := RummyConfig{
		plugins: plugins,
		base:    initLocalConfig(),
	}

	app.installPlugins()
}

func (rc RummyConfig) installPlugins() {
	for _, plugin := range rc.plugins {
		plugin.Install()
	}
}
