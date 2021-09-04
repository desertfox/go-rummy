package main

import (
	"os"

	flag "github.com/spf13/pflag"

	"github.com/go-rummy/pkg"
)

var (
	installList []string
	overwrite   bool
)

func init() {
	flag.StringSliceVar(&installList, "i", []string{"all"}, "Install")
	flag.BoolVar(&overwrite, "o", false, "Source path/dir for dot files")
}

func main() {
	flag.Parse()

	home := os.Getenv("HOME")

	r := &rummy.Rummy{
		InstallList: installList,
		DestDir:     home,
		Overwrite:   overwrite,
	}

	err := r.Go()
	if err != nil {
		panic(err)
	}
}
