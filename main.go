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
	list            = make(map[string]rummy.Installer)
)

func init() {
	flag.StringSliceVar(&install, "i", []string{"all"}, "Install")
	flag.StringVar(&dotFiles, "df", defaultDotFiles, "Source path/dir for dot files")
}

func main() {
	flag.Parse()

	buildList()

	plugins := buildPlugins()

	rummy.Go(plugins)

}

func buildList() {
	buildDotfiles()

	list["vim"] = p.NewVimPlugin(dotFiles)
	list["bash"] = p.NewBashPlugin(dotFiles)
}

func buildDotfiles() {
	if dotFiles == defaultDotFiles {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		dotFiles = filepath.Join(wd, defaultDotFiles)
	}
}

func buildPlugins() *[]rummy.Installer {
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
