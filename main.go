package main

import (
	"os"

	flag "github.com/spf13/pflag"

	"github.com/go-rummy/pkg"
	p "github.com/go-rummy/pkg/plugins"
)

var (
	install   []string
	overwrite bool
)

func init() {
	flag.StringSliceVar(&install, "i", []string{"all"}, "Install")
	flag.BoolVar(&overwrite, "o", false, "Source path/dir for dot files")
}

func main() {
	flag.Parse()

	home := os.Getenv("HOME")

	ap := buildAvailablePluginList(home, overwrite)

	plugins := selectPlugins(ap)

	rummy.Go(plugins)

}

func buildAvailablePluginList(destDir string, overwrite bool) map[string]rummy.Installer {
	list := make(map[string]rummy.Installer)

	list["git"] = p.NewGitPlugin(destDir, overwrite)
	list["vim"] = p.NewVimPlugin(destDir, overwrite)
	list["zsh"] = p.NewZshPlugin(destDir, overwrite)
	list["bash"] = p.NewBashPlugin(destDir, overwrite)

	return list
}

func selectPlugins(list map[string]rummy.Installer) *[]rummy.Installer {
	plugins := make([]rummy.Installer, 0, len(list))

	if install[0] == "all" {
		for _, v := range list {
			plugins = append(plugins, v)
		}

	} else {
		for _, v := range install {
			if _, exists := list[v]; exists {
				plugins = append(plugins, list[v])
			} else {
				panic("No plugin found for " + v)
			}
		}
	}

	return &plugins
}
