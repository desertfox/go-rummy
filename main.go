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

	if dotFiles == defaultDotFiles {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		dotFiles = filepath.Join(wd, defaultDotFiles)
	}

	list := make(map[string]rummy.Installer)
	list["vim"] = p.NewVimPlugin(dotFiles)
	list["bash"] = p.NewBashPlugin(dotFiles)

	if install[0] == "all" {
		plugins := make([]rummy.Installer, 0, len(list))
		for _, v := range list {
			plugins = append(plugins, v)
		}

		rummy.Go(plugins)
	} else {
		plugins := make([]rummy.Installer, 0, len(install))
		for _, v := range install {
			plugins = append(plugins, list[v])
		}
		rummy.Go(plugins)
	}

}
