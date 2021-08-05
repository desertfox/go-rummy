package main

import (
	"os"

	"github.com/go-rummy/pkg"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	rummy.GoRummy(wd)
}
