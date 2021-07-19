package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var dotfilePath string

func main() {

	var overwriteVimrc = flag.Bool("ovimrc", false, "Overwrite .vimrc file")

	flag.Parse()

	installBashConfigs()

	installVimrc(overwriteVimrc)

	installVimPlug()
}

func init() {
	//Set cwd, make sure dot-files exists
	fmt.Println("Init")
	//make configurable
	const Dotfiles = "dot-files"

	cwd, err := os.Getwd()
	check(err)

	dotfilePath = strings.Join([]string{cwd, Dotfiles}, "/")
}

func installBashConfigs() {
	// Move alias and profile
	fmt.Println("installBashConfig")
}

func initConfigFile(source string, dest string, overwrite *bool) {
	source = strings.Join([]string{dotfilePath, source}, "/")

	sourceFile, err := os.Open(source)
	check(err)
	defer sourceFile.Close()

	dest = strings.Join([]string{os.Getenv("HOME"), dest}, "/")
	if _, err := os.Stat(dest); err == nil {
		//Fix vimrc name
		fmt.Println("%v file already exists", dest)

		if *overwrite == false {
			fmt.Println("%v overwrite not set, bailing.", dest)
			return
		}

		fmt.Println("%v overwrite set, vimrc will be overwritten.", dest)
	}

	destinationFile, err := os.Create(dest)
	check(err)
	defer destinationFile.Close()

	bytesCopied, err := io.Copy(destinationFile, sourceFile)
	check(err)

	fmt.Println("Installed %v bytes:%v", dest, bytesCopied)

	return
}

func installVimrc(ovimrc *bool) {
	const Vimrc = ".vimrc"
	sourceVimrc := strings.Join([]string{dotfilePath, Vimrc}, "/")

	source, err := os.Open(sourceVimrc)
	check(err)
	defer source.Close()

	dst := strings.Join([]string{os.Getenv("HOME"), Vimrc}, "/")
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
	//Move to config
	const (
		VimPlugUrl      = "https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim"
		VimPlugDestFile = "/.vim/autoload/plug.vim"
	)
	var vimAutoloadPath = strings.Join([]string{os.Getenv("HOME"), VimPlugDestFile}, "/")

	resp, err := http.Get(VimPlugUrl)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	if _, err := os.Stat(vimAutoloadPath); err == nil {
		fmt.Printf("vim plug.vim file already exists\n")
		return
	}

	err = ioutil.WriteFile(vimAutoloadPath, body, 0644)
	check(err)

	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
