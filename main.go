package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	var overwriteVimrc = flag.Bool("ovimrc", false, "Overwrite .vimrc file")

	flag.Parse()

	installVimrc(overwriteVimrc)
	installVimPlug()
}

func installVimrc(ovimrc *bool) {
	sourceVimrc, err := os.Getwd()
	check(err)
	sourceVimrc = sourceVimrc + "/dot-files/.vimrc"

	source, err := os.Open(sourceVimrc)
	check(err)
	defer source.Close()

	dst := os.Getenv("HOME") + "/.vimrc"
	if _, err := os.Stat(dst); err == nil {
		fmt.Println("vimrc file already exists")

		if *ovimrc == false {
			fmt.Println("overwrite not set, bailing.")
			return
		}

		fmt.Println("overwrite set, vimrc will be overwritten.")
	}

	destination, err := os.Create(dst)
	check(err)
	defer destination.Close()

	_, err = io.Copy(destination, source)
	check(err)

	fmt.Printf("Installed vimrc\n")

	return
}

func installVimPlug() {
	const VimPlugUrl = "https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim"
	var VimAutoloadPath = os.Getenv("HOME") + "/.vim/autoload/plug.vim"

	resp, err := http.Get(VimPlugUrl)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	if _, err := os.Stat(VimAutoloadPath); err == nil {
		fmt.Printf("vim plug.vim file already exists\n")
		return
	}

	err = ioutil.WriteFile(VimAutoloadPath, body, 0644)
	check(err)

	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
