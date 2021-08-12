package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/go-rummy/pkg"
)

var (
	install         string
	dotFiles        string
	defaultDotFiles = "dot-files"
)

func init() {
	flag.StringVar(&install, "i", "all", "Install")
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

	if install == "all" {
		rummy.Go(dotFiles, []string{"all"})
	}

}
