package main

import (
	"os"
	"path/filepath"

	"github.com/go-rummy/pkg"
	p "github.com/go-rummy/pkg/plugins"
	flag "github.com/spf13/pflag"
)

var (
	install         []string
	dotFiles        string
	defaultDotFiles = "dot-files"
)

func init() {
	flag.StringSliceVar(&install, "i", []string{"all"}, "Install")
	flag.StringVar(&dotFiles, "df", defaultDotFiles, "Source path/dir for dot files")
}

func main() {
	flag.Parse()

	sd := buildSourceDir(dotFiles, defaultDotFiles)

	ap := buildAvailablePluginList(sd)

	plugins := selectPlugins(ap)

	rummy.Go(plugins)

}

func buildSourceDir(df, ddf string) string {
	if df == ddf {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		return filepath.Join(wd, ddf)
	}

	return df
}

func buildAvailablePluginList(sourceDir string) map[string]rummy.Installer {
	list := make(map[string]rummy.Installer)
	list["vim"] = p.NewVimPlugin(sourceDir, os.Getenv("HOME"))
	list["bash"] = p.NewBashPlugin(sourceDir, os.Getenv("HOME"))

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
