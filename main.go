package main

import (
	"flag"
	"os"

	"github.com/go-rummy/pkg"
)

type installList []string

var (
	appList   installList
	overwrite bool
)

func (i *installList) String() string {
	return ""
}

func (i *installList) Set(app string) error {
	*i = append(*i, app)
	return nil
}

func init() {
	flag.Var(&appList, "i", "Install")
	flag.BoolVar(&overwrite, "o", false, "Overwrite")
}

func main() {
	flag.Parse()

	home := os.Getenv("HOME")

	if appList == nil {
		appList = installList{"all"}
	}

	r := &rummy.Rummy{
		InstallList: appList,
		DestDir:     home,
		Overwrite:   overwrite,
	}

	if err := r.Go(); err != nil {
		panic(err)
	}
}
