package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	_, err := installVim()
	if err != nil {
		panic(err.Error())
	}
}

func installVim() (int64, error) {
	workdir, err := os.Getwd()
	if err != nil {
		return 0, err
	}
	sourceVimrc := workdir + "/dot-files/.vimrc"

	source, err := os.Open(sourceVimrc)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	dst := os.Getenv("HOME") + "/.vimrc"
	if _, err := os.Stat(dst); err == nil {
		fmt.Printf("vimrc file already exists\n")
		return 0, err
	}

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	if err != nil {
		return 0, err
	}
	fmt.Printf("Installed vimrc\n")

	return nBytes, err
}
